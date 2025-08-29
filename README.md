# Synapse-
Distributed Task Management System

curl -X POST http://localhost:8081/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Написать API",
    "description": "Сделать CRUD для задач",
  }'
