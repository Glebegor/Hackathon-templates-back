
import express, { Request, Response } from 'express';
import getConfig  from './config';
import { IConfig } from '../core/commonServer/config';

class Application {
    public app: express.Application;
    public config: IConfig;
    
    constructor(envType: string) {
        this.config = getConfig(envType);
        this.app = express();
    }

    public run() {
        this.app.listen(this.config.SERVER.PORT, () => {
            console.log(`Server running on port ${this.config.SERVER.PORT}`);
        });
    }
}

export default Application;
