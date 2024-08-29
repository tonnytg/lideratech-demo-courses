#!/bin/bash

echo "Instalando os pacotes"

pip install Flask

pip install > requirements.txt

gcloud app deploy

gcloud app browse
