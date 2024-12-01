

import { Router } from 'express';
import Application from './bootstrap/application';
import { newRouterMain } from './routes/routerMain';


// get args from npm start command
const args = process.argv.slice(2);

// In the base, you have 3 types of environments: dev, test, and prod
const envType = args[0] || 'dev';


const app = new Application(envType);

app.initialize()
    .catch((error) => {
        console.error('Error initializing application', error);
        process.exit(1);
    });

const router: Router = newRouterMain(app.config, app.dbClient).run();

app.app.use('/api/v2', router);
app.run();