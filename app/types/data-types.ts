

export type InflateUI = {
    spends: Spending[]
    currency_rate: CurrencyRate

}


export interface Spending {
    Id: number
    Name: string
    Type: string
    Amount: number
    Currency: string
    Description: string
    Time: string
}

export interface CurrencyRate {
    base?: string
    timestamp?: number
    rates?: Record<string, number>
}
