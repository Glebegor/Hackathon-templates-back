
import express, { Request, Response } from 'express';
import getConfig  from './config';
import { IConfig } from '../core/commonServer/config';
import { PrismaClient } from '@prisma/client';

class Application {
    public app: express.Application;
    public config: IConfig;
    public dbClient: PrismaClient;
    
    constructor(envType: string) {
        this.config = getConfig(envType);
        this.dbClient = new PrismaClient();   
             
        this.app = express();
        this.app.use(express.json());

    }

    private pingDb(): Promise<string> {
        return this.dbClient.$connect()
            .then(() => {
                console.log('Database connected');
                return 'Database connected';
            })
            .catch((error: Error) => {
                console.error('Error connecting to database', error);
                return error.message;
            });
    }

    public async initialize(): Promise<void> {
        try {
            await this.pingDb();

        } catch (error) {
            console.error('Error initializing application', error);
            process.exit(1);
        }
    }


    public run(): void {
        this.app.listen(this.config.SERVER.PORT, () => {
            console.log(`Server running on port ${this.config.SERVER.PORT}`);
        });
    }
}

export default Application;
