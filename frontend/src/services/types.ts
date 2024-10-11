export interface ExchangeInfo {
  supported_margin_types: string[];
  leaderboard_periods: SupportedPeriods;
  leaderboard_fields: string[];
}

export interface InfoResponse {
  exchange_info: { [key: string]: ExchangeInfo };
}

export interface SupportedPeriods {
  fixed_periods: string[];
  custom_periods?: CustomPeriods;
}

export interface CustomPeriods {
  min_days: number;
  max_days: number;
}

export interface APIError {
  code: number;
  message: string;
  details?: string;
}

export interface Trader {
  user_id: string;
  period_pnl_percent?: number;
  period_pnl_absolute?: number;
  total_trades?: number;
  wins?: number;
  volume?: number;
  avg_win?: number;
  avg_loss?: number;
}

export interface FuturesResponse {
  trader: string;
  platform: string;
  updated_positions: FuturesPosition[];
  detected_changes: FuturesDelta[];
  is_initial_positions: boolean;
}

export interface FuturesPosition {
  market: string;
  entry_price: number;
  current_price: number;
  status: Status;
  direction: Direction;
  margin_type: MarginType;
  size_usd: number;
  size_tokens: number;
  unrealized_pnl: number;
  leverage_amount?: number;
  collateral_token_amount?: number;
  collateral_token_amount_usd?: number;
  collateral_token?: string;
  health_ratio?: number;
  cross_margin_share?: number;
  free_collateral?: number;
}

export interface FuturesDelta {
  market: string;
  status: Status;
  direction: Direction;
  size_tokens_delta: number;
  size_usd_delta: number;
  leverage_delta: number;
  collateral_token?: string;
  collateral_token_amount_delta?: number;
}

export enum Direction {
  Long = "long",
  Short = "short",
}

export enum MarginType {
  Isolated = "isolated",
  Cross = "cross",
}

export enum Status {
  Open = "open",
  Closed = "closed",
  Liquidated = "liquidated",
}
