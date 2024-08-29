# Subindo um site com Google Cloud Storage (GCS)

O Bucket permite que você suba um site e disponibilize ele para o mundo inteiro com se fosse uma CDN.<br />
Você pode usar esse recurso para validar idéias e demos para os seus clientes de forma rápida.


### Video

[![Assista ao vídeo no YouTube](https://img.youtube.com/vi/SSafDMkiBJQ/0.jpg)](https://www.youtube.com/watch?v=SSafDMkiBJQ&t)



### How to

Se autentique na GCP

    gcloud auth login

Defina o projeto padrão a ser utilizado

    gcloud config set project YOUR_PROJECT_ID


Crie seu Bucket na GCP na região que você quer

    gsutil mb -l REGION gs://YOUR_BUCKET_NAME


Copie os arquivos para dentro do Bucket

    gsutil cp index.html gs://meu-bucket-exemplo


Defina as permissões para o arquivo ser público

    gsutil acl ch -u AllUsers:R gs://meu-bucket-exemplo/index.html


Realize o teste da conexão para validar se o site está em pé

    curl -v https://storage.googleapis.com/meu-bucket-exemplo/index.html

