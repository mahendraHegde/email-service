gqlgen generate
heroku config:set $(cat .env | sed '/^$/d; /#[[:print:]]*$/d')