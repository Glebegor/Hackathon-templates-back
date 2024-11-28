interface IConfig {
    SERVER: {
        PORT: number,
        SECRET: string | undefined
    },
    DB: {
        HOST: string,
        PORT: number,
        USER: string,
        PASSWORD: string | undefined,
        DATABASE: string
    }
}

export { IConfig };