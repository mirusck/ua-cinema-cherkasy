# Ukrainian movies in Cherkasy

---

The goal of the project is to create easy and useful notifier 
about new Ukrainian movies that just came out in local cinemas.    

## Environment variables

```
# URL of movie theatre to parse
export CINEMA_URL=http://ukraina.ck.ua
# AWS configs
export AWS_REGION=eu-central-1
export AWS_ACCESS_KEY_ID=your_aws_key
export AWS_SECRET_ACCESS_KEY=your_aws_secret_key
# Facebook Messenger secrets
export PAGE_ACCESS_TOKEN=your_page_access_token
export VERIFY_TOKEN=some_random_verify_token
```

## Tests
Run `ginkgo -r`
