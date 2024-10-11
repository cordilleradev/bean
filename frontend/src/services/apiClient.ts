import { APIError, InfoResponse, Trader, FuturesResponse } from "./types";
export class APIClient {
  private baseUrl: string;

  constructor(baseUrl: string) {
    this.baseUrl = baseUrl;
  }

  private async fetchJson<T>(
    endpoint: string,
    options: RequestInit = {}
  ): Promise<T> {
    const response = await fetch(`${this.baseUrl}${endpoint}`, options);
    if (!response.ok) {
      const errorData: APIError = await response.json();
      throw new Error(`API Error: ${errorData.message}`);
    }
    return response.json() as Promise<T>;
  }

  async getInfo(): Promise<InfoResponse> {
    return this.fetchJson<InfoResponse>("/info");
  }

  async getLiveLeaderboard(
    exchange: string,
    params: {
      limit?: number;
      order?: "asc" | "desc";
      period?: string;
      sort_by?: string;
    }
  ): Promise<Trader[]> {
    const queryParams = new URLSearchParams(
      params as Record<string, string>
    ).toString();
    return this.fetchJson<Trader[]>(
      `/live_leaderboard/${exchange}?${queryParams}`
    );
  }

  async getPositions(
    exchange: string,
    userId: string
  ): Promise<FuturesResponse> {
    return this.fetchJson<FuturesResponse>(`/positions/${exchange}/${userId}`);
  }
}
