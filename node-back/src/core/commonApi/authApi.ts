
// Login
interface ILoginRequest {
    username: string,
    email: string
    password: string
}

interface ILoginResponse {
    accessToken: string,
    refreshToken: string
}

// Register
interface IRegisterRequest {
    username: string,
    email: string,
    password: string
}

// Refresh
interface IRefreshTokenRequest {
    refreshToken: string
}

interface IRefreshTokenResponse {
    accessToken: string
    refreshToken: string
}

export { ILoginRequest, ILoginResponse, IRegisterRequest, IRefreshTokenRequest, IRefreshTokenResponse };