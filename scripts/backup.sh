#!/bin/bash
docker exec -u postgres pgsql pg_dump -Fc owl > db.dump
