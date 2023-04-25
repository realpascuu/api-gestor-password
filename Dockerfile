# Dockerfile to create postgres default environment
FROM postgres:14.7-alpine
ENV POSTGRES_PASSWORD docker
ENV POSTGRES_DB gestorpassword
