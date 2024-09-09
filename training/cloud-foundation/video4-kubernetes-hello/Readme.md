# Kubernetes and Cloud

As nuvens públicas criaram suas próprias soluções de Kubernetes para simplificar a adoção dessa tecnologia, oferecendo clusters gerenciados que eliminam parte da complexidade de administração e operação. Entre essas soluções, o Google Kubernetes Engine (GKE) se destaca pela integração nativa com o Google Cloud, permitindo um gerenciamento mais eficiente e seguro dos seus clusters. No entanto, mesmo com essas facilidades, é essencial ter um bom conhecimento da estrutura do Kubernetes para aproveitar ao máximo essas plataformas.

## Control Plane

O Control Plane é o coração do Kubernetes. Ele gerencia o estado desejado do cluster, assegurando que as aplicações rodando no ambiente estejam sempre de acordo com as definições especificadas. O Control Plane é composto por vários componentes essenciais:

- **API Server**: Serve como a interface principal para a comunicação com o cluster. Todas as operações do Kubernetes passam por aqui.
- **Etcd**: Um banco de dados chave-valor altamente disponível e distribuído que armazena todos os dados de configuração do cluster.
- **Scheduler**: Responsável por atribuir pods a nós disponíveis, garantindo que os recursos sejam utilizados da forma mais eficiente possível.
- **Controller Manager**: Monitora o estado do cluster e garante que ele corresponda ao estado desejado, realizando ajustes automáticos quando necessário.

No GKE, o Control Plane é gerenciado pelo Google, eliminando a necessidade de configuração e manutenção manual, e permitindo que você se concentre mais nas suas aplicações e menos na infraestrutura.

## Node Pools

No GKE, os nós que compõem o cluster são organizados em pools, chamados de Node Pools. Cada pool pode ter uma configuração diferente, como tipos de máquina, versões de sistema operacional, e escopo de tolerância a falhas. Isso oferece flexibilidade na alocação de recursos para diferentes tipos de workloads, otimizando custos e desempenho.

- **Auto-Scaling**: O GKE oferece autoescalamento para os Node Pools, ajustando automaticamente o número de nós de acordo com a carga de trabalho atual. Isso garante que seus aplicativos tenham os recursos necessários em momentos de pico e economizem custos durante períodos de baixa utilização.

## Workloads

No Kubernetes, as workloads são definidas como unidades de computação. O GKE oferece suporte completo para diversos tipos de workloads, incluindo:

- **Pods**: A menor unidade do Kubernetes, representando uma ou mais contêineres que compartilham o mesmo espaço de rede e armazenamento.
- **Deployments**: Gerenciam a escalabilidade e o ciclo de vida dos pods, garantindo alta disponibilidade e atualizações contínuas.
- **StatefulSets**: Úteis para aplicações que necessitam de um armazenamento persistente, garantindo que cada pod tenha um identificador único e fixo.
- **DaemonSets**: Garantem que uma cópia de um pod seja executada em cada nó do cluster.

## Networking

A rede no Kubernetes é complexa, mas GKE facilita sua implementação e gerenciamento:

- **Service**: Um objeto Kubernetes que expõe uma aplicação em execução como um serviço de rede. Ele pode ser do tipo `ClusterIP`, `NodePort`, ou `LoadBalancer`.
- **Ingress**: GKE suporta Ingress para rotear tráfego externo para os serviços dentro do cluster, fornecendo load balancing e SSL/TLS nativos.

## Segurança

A segurança é uma prioridade no GKE, que oferece recursos integrados para proteger seu cluster:

- **Autenticação e Autorização**: GKE utiliza as contas de serviço do Google Cloud para autenticação, integrando-se ao IAM (Identity and Access Management) para controle de acesso.
- **Network Policies**: Permitem controlar o tráfego entre pods, adicionando uma camada extra de segurança.
- **Node Security**: GKE pode provisionar nós com o Kubernetes Sandbox (GVisor), oferecendo um ambiente isolado para executar containers, aumentando a segurança contra ataques.

