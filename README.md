### Deploy to Heroku CircleCI step

To make it working

-   generate a new SSH private-public key pair
-   register the public key to your Heroku account
-   add the private key to CircleCI, to allow it publishing

#### Useful guides:

-   https://devcenter.heroku.com/articles/keys
-   https://circleci.com/docs/2.0/add-ssh-key/
-   https://www.ssh.com/ssh/public-key-authentication

## todo

-   split 'fmt', 'vet' and 'test' into separate steps
-   add env variable for storing cache version
-   remove '-f' force flag from heroku push
