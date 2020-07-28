# -*- mode: Python -*-

# Please ensure to change NODE_VERSION in .env file as well.
node_version = "12"
docker_compose('docker-compose.yml')

# docker_build(
#   # Image name - must match the image in the docker-compose file
#   'srijanlabs/node:' + node_version + '-buster-local',
#   '.',
#   build_args={'NODE_VERSION': node_version},
#   dockerfile='Dockerfile.dev',
#   live_update = [
#     # Sync local files into the container.
#     sync('.', '/app'),
#     run('npm i', trigger='package.json'),
#     restart_container()
#   ])
