1. Install Google Cloud SDK:

   - Make sure you have the Google Cloud SDK (gcloud) installed on your local machine. You can install it by following the instructions here: https://cloud.google.com/sdk/docs/install

2. Create a project on Google Cloud Console: (Syns Platform)

   - Go to the Google Cloud Console (https://console.cloud.google.com/) and create a new project or select an existing one. Take note of the Project ID.

3. Initialize the SDK:

   - Open your terminal/command prompt and run gcloud init to configure the SDK with your Google Cloud account and set the default project to the one you created or selected earlier.

4. Prepare Dockerfile

5. Build the container image:

   - Run the following command to build the container image and store it in Google Container Registry (gcr.io):

```
    make gcloud-build
```

- notice: find the image in `Cloud Build` section

6. For the particular setup in .env file, the first deployment using google cloud sdk is quite problematic. Instead of using google cloud sdk, deploy the image using google cloud console.

   a. Go to the Google Cloud Console: https://console.cloud.google.com/

   b. Make sure you have the correct project selected. If not, click on the project dropdown at the top right corner and select the desired project from the list.

   c. In the left-side menu, click on "Cloud Run" under the "Serverless" section.

   d. Click on the "Create service" button => fill out the necessary information => REMEMBER to check on `Allow unauthenticated invocations
Check this if you are creating a public API or website.`

   e. At the end of the console, click the drop down "Container, Networking, Security" => add all the environment variables and REMEMBER to include GIN_MODE=release variable

   f. Click Create!

7. Now the service is successfully created, if there are any changes in the codes. Just run

```
    make gcloud-deploy
```

this will rebuild, push the docker image and redeploy the container.
