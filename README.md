# other_than_gcp
GCP内から、G SuiteとかFirebaseを叩くぞい

## Local

`gcloud auth application-default login` だけだとGCPのScopeしか取っていないので、G Suiteなどを実行する時は明示的にScopeを設定する。

```
gcloud auth application-default login --scopes "https://www.googleapis.com/auth/spreadsheets"
```

## GCP

### Spreadsheets

Service Accountが対象のSpreadsheetsにアクセスする権限を持っている必要がある。