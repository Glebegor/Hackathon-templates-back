
interface IUser {
    id: number | undefined,
    username: string,
    email: string,
    passwordHash: string
}

export { IUser };