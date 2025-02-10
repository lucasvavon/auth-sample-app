# Build the application from source
FROM golang:1.23.4-alpine3.21

WORKDIR /app

# Copier les fichiers go.mod et go.sum et télécharger les dépendances
COPY go.mod go.sum ./
RUN go mod download

# Copier le reste du code source et compiler l'application
COPY . .
RUN go build -o main .

# Image finale
FROM alpine:latest

WORKDIR /root/

# Installer les dépendances nécessaires
RUN apk --no-cache add ca-certificates

# Copier l'exécutable depuis l'étape de build
COPY --from=builder /app/main .

# Définir les variables d'environnement pour PostgreSQL
ENV DATABASE_URL="postgresql://admin:admin@db:5432/myapp"

# Exposer le port si nécessaire
EXPOSE 1323

# Lancer l'application
CMD ["./main"]