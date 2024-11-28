// get args from npm start command
const args = process.argv.slice(2);

// In the base, you have 3 types of environments: dev, test, and prod
const envType = args[0] || 'dev';

import Application from './bootstrap/application';
const app = new Application(envType);

app.initialize()
    .then(() => app.run())
    .catch((error) => {
        console.error('Error initializing application', error);
        process.exit(1);
    });
