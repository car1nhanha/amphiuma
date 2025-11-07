package sina

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/car1nhanha/go-blog/internal/db"
	"github.com/car1nhanha/go-blog/internal/model"
)

func getAllFilesFromGithub(user string) (model.IGithubSearchFile, error) {
	var files model.IGithubSearchFile

	url := "https://api.github.com/search/code?q=extension:go+user:" + user
	log.Printf("\033[35mSina - getAllFilesFromGithub %s\033[0m\n\n", url)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return files, err
	}

	request.Header.Set("Authorization", os.Getenv("TOKEN"))

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Printf("erro ao consultar os arquivos de %s %v", user, err)
		return files, err
	}
	// request.Body.Close()

	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		log.Printf("erro na requisição dos repositórios de %s: %s", user, string(body))
		return files, fmt.Errorf("status %d", response.StatusCode)
	}

	json.NewDecoder(response.Body).Decode(&files)
	return files, nil
}

func FindFilesFromDB(user string) ([]model.File, error) {
	log.Printf("\033[35msina - FindFilesFromDB %s", user)
	// recebe um userId
	// busca todos os dados do userId no banco com ordenação por data
	out, err := db.Client.Query(context.Background(), &dynamodb.QueryInput{
		TableName:              aws.String("files"),
		KeyConditionExpression: aws.String("#u = :owner_login"),
		ExpressionAttributeNames: map[string]string{
			"#u": "owner_login",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":owner_login": &types.AttributeValueMemberS{Value: user},
		},
	})
	if err != nil {
		return nil, err
	}

	var files []model.File
	if err := attributevalue.UnmarshalListOfMaps(out.Items, &files); err != nil {
		return nil, err
	}
	log.Printf("rose - FindFilesFromDB %s - return", user)
	return files, nil
}

// comparar os arquivos que estão no banco com os que chegou da api
func compareFile(filesGithub *model.IGithubSearchFile, filesDB []model.File) (present, notPresent, noLongerExists []string) {
	// 🔹 1. Mapa rápido dos arquivos do banco para busca O(1)
	dbMap := make(map[string]struct{}, len(filesDB))
	for _, f := range filesDB {
		key := f.RepoId + ":" + f.Path // chave única por repositório + caminho
		dbMap[key] = struct{}{}
	}

	// 🔹 2. Verifica os arquivos do GitHub (se estão ou não no banco)
	seen := make(map[string]struct{}, len(filesGithub.Items))
	for _, g := range filesGithub.Items {
		key := fmt.Sprintf("%d:%s", g.Repository.ID, g.Path)
		seen[key] = struct{}{}

		if _, exists := dbMap[key]; exists {
			present = append(present, key)
		} else {
			notPresent = append(notPresent, key)
		}
	}

	// 🔹 3. Verifica os arquivos do banco que não estão mais no GitHub
	for _, f := range filesDB {
		key := f.RepoId + ":" + f.Path
		if _, stillThere := seen[key]; !stillThere {
			noLongerExists = append(noLongerExists, key)
		}
	}

	log.Printf("present: %v", present)
	log.Printf("not present: %v", notPresent)
	log.Printf("no longer exists: %v", noLongerExists)

	return present, notPresent, noLongerExists
}

// SyncUserFiles sincroniza os arquivos de um usuário com o GitHub
func SyncUserFiles(user string) error {
	log.Printf("\033[34mbuscando dados do github")
	// 1. Buscar arquivos do GitHub
	filesGithub, err := getAllFilesFromGithub(user)
	if err != nil {
		return fmt.Errorf("erro ao buscar arquivos do GitHub: %w", err)
	}
	log.Printf("encontrados => %d", len(filesGithub.Items))
	// 2. Buscar arquivos do banco
	filesDB, err := FindFilesFromDB(user)
	if err != nil {
		// Se não encontrar arquivos, considera como lista vazia
		filesDB = []model.File{}
	}

	// 3. Comparar arquivos
	_, notPresent, noLongerExists := compareFile(&filesGithub, filesDB)

	// 4. Adicionar arquivos faltantes
	if len(notPresent) > 0 {
		log.Printf("Adicionando %d arquivos novos...", len(notPresent))
		if err := addMissingFiles(&filesGithub, notPresent, user); err != nil {
			return fmt.Errorf("erro ao adicionar arquivos: %w", err)
		}
	}

	// 5. Remover arquivos que não existem mais
	if len(noLongerExists) > 0 {
		log.Printf("Removendo %d arquivos antigos...", len(noLongerExists))
		if err := removeDeletedFiles(filesDB, noLongerExists); err != nil {
			return fmt.Errorf("erro ao remover arquivos: %w", err)
		}
	}

	log.Printf("✅ Sincronização concluída para %s", user)
	return nil
}

// addMissingFiles adiciona os arquivos que estão no GitHub mas não no banco
func addMissingFiles(filesGithub *model.IGithubSearchFile, notPresent []string, user string) error {
	// Criar mapa para busca rápida
	keysToAdd := make(map[string]struct{})
	for _, key := range notPresent {
		keysToAdd[key] = struct{}{}
	}

	ctx := context.Background()
	errCount := 0

	for _, item := range filesGithub.Items {
		key := fmt.Sprintf("%d:%s", item.Repository.ID, item.Path)

		if _, shouldAdd := keysToAdd[key]; !shouldAdd {
			continue
		}

		// Criar o modelo File
		file := model.File{
			Name:       item.Name,
			Path:       item.Path,
			URL:        item.HTMLURL,
			RepoId:     strconv.Itoa(item.Repository.ID),
			OwnerLogin: user,
			OwnerId:    strconv.Itoa(item.Repository.Owner.ID),
			CreatedAt:  time.Now(),
		}

		// Converter para DynamoDB attributes
		av, err := attributevalue.MarshalMap(file)
		if err != nil {
			log.Printf("❌ Erro ao serializar arquivo %s: %v", key, err)
			errCount++
			continue
		}

		// Inserir no DynamoDB
		_, err = db.Client.PutItem(ctx, &dynamodb.PutItemInput{
			TableName: aws.String("files"),
			Item:      av,
		})
		if err != nil {
			log.Printf("❌ Erro ao adicionar arquivo %s: %v", key, err)
			errCount++
			continue
		}

		log.Printf("✅ Arquivo adicionado: %s", key)
	}

	if errCount > 0 {
		return fmt.Errorf("falha ao adicionar %d arquivo(s)", errCount)
	}

	return nil
}

// removeDeletedFiles remove os arquivos que não existem mais no GitHub
func removeDeletedFiles(filesDB []model.File, noLongerExists []string) error {
	// Criar mapa para busca rápida
	keysToRemove := make(map[string]struct{})
	for _, key := range noLongerExists {
		keysToRemove[key] = struct{}{}
	}

	ctx := context.Background()
	errCount := 0

	for _, file := range filesDB {
		key := file.RepoId + ":" + file.Path

		if _, shouldRemove := keysToRemove[key]; !shouldRemove {
			continue
		}

		// IMPORTANTE: Como sua tabela só tem HASH key (owner_login),
		// precisamos incluir TODOS os atributos que formam a chave primária
		// Se houver uma RANGE key definida, você precisa incluí-la aqui também
		_, err := db.Client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
			TableName: aws.String("files"),
			Key: map[string]types.AttributeValue{
				"owner_login": &types.AttributeValueMemberS{Value: file.OwnerLogin},
				// ⚠️ Se sua tabela tiver RANGE key, adicione aqui:
				// "path": &types.AttributeValueMemberS{Value: file.Path},
			},
		})
		if err != nil {
			log.Printf("❌ Erro ao remover arquivo %s: %v", key, err)
			errCount++
			continue
		}

		log.Printf("🗑️  Arquivo removido: %s", key)
	}

	if errCount > 0 {
		return fmt.Errorf("falha ao remover %d arquivo(s)", errCount)
	}

	return nil
}
