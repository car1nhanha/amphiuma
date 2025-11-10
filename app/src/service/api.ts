type IInfo = {
  login: string;
  avatar_url: string;
  html_url: string;
  name: string;
  public_repos: number;
  followers: number;
  following: number;
};

export class Api {
  private api_url = import.meta.env.VITE_API_BACKEND;

  async getInfo() {
    fetch(`${this.api_url}`)
      .then((e) => {
        return e.body;
      })
      .catch((err) => {
        console.log();
      });
  }
}
