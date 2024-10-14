# thesaurus

## Description

This is a basic React application for a tiny thesaurus that can only provide one suggestion for any two-letter words. The application allows users to input a two-letter word and get a suggestion for it.

## How to Run the Project

1. Clone the repository:
   ```
   git clone https://github.com/johnpfeiffer/thesaurus.git
   ```
2. Navigate to the project directory:
   ```
   cd thesaurus
   ```
3. Install the dependencies:
   ```
   npm install
   ```
4. Start the development server:
   ```
   npm start
   ```
5. Open your browser and go to `http://localhost:3000` to see the application.

## How to Deploy the Project using GitHub Actions

1. Create a new file in the `.github/workflows` directory called `deploy.yml`.
2. Add the following content to the `deploy.yml` file:
   ```yaml
   name: Deploy

   on:
     push:
       branches:
         - main

   jobs:
     build:
       runs-on: ubuntu-latest

       steps:
         - name: Checkout code
           uses: actions/checkout@v2

         - name: Set up Node.js
           uses: actions/setup-node@v2
           with:
             node-version: '14'

         - name: Install dependencies
           run: npm install

         - name: Build the app
           run: npm run build

         - name: Deploy to GitHub Pages
           uses: peaceiris/actions-gh-pages@v3
           with:
             github_token: ${{ secrets.GITHUB_TOKEN }}
             publish_dir: ./build
   ```
3. Commit and push your changes to the `main` branch.
4. The GitHub Actions workflow will automatically build and deploy your app to GitHub Pages.
