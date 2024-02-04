# How to start

    docker-compose up -d

# How to get in the mysql and start the container

    # Start container
        docker-compose exec mysql bash
        mysql -uroot -p goexpert

    # Create the module
        go mod init github.com/osniantonio/goexpert/12/5
