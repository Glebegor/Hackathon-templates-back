// get args from npm start command
const args = process.argv.slice(2);

// In the base, you have 3 types of environments: dev, test, and prod
const envType = args[0] || 'dev';

import { Router } from 'express';
import Application from './bootstrap/application';
import { newControllerMain } from './controllers/controllerMain';
const app = new Application(envType);

app.initialize()
    .catch((error) => {
        console.error('Error initializing application', error);
        process.exit(1);
    });

// Generate main controller
const controller: Router = newControllerMain(app.config, app.dbClient);

app.app.use('/api/v1', controller);

app.run();