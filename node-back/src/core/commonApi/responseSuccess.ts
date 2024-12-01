

class ResponseSuccess {
    public status: number;
    public message: string;
    public data: any;

    constructor(status: number, message: string, data: any) {
        this.status = status;
        this.message = message;
        this.data = data;
    }

    public send(res: any): void {
        res.status(this.status).send({
            message: this.message,
            data: this.data
        });
    }
}

export { ResponseSuccess };