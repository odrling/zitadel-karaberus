# Zitadel Karaberus

This can be used as a simple oidc server for development on Karaberus.

```sh
go run github.com/odrling/zitadel-karaberus/server@master
```

This will start the server on port 9998.
Username is `admin` or `user` and password `verysecure`.

Example configuration for karaberus:
```
export KARABERUS_OIDC_ISSUER=http://localhost:9998/
export KARABERUS_OIDC_KEY_ID=karaberus
export KARABERUS_OIDC_CLIENT_ID=web
export KARABERUS_OIDC_CLIENT_SECRET=secret
export KARABERUS_OIDC_GROUPS_CLAIM=groups
export KARABERUS_OIDC_ADMIN_GROUP=admin
export KARABERUS_OIDC_JWT_SIGN_KEY="$(openssl rand -hex 32)"
```

When using the UI dev server you should also add this variable:
```
export KARABERUS_LISTEN_BASE_URL=http://localhost:5173
```
