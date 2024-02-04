# How to start

    docker-compose up -d

# How to get in the mysql and start the container

    # Start container
        docker-compose exec mysql bash
        mysql -uroot -p goexpert

    # Create the table
        msql>
            create table products (id varchar(255), name varchar(80), price decimal(10, 2), primary key(id));

    # Create the module
        go mod init github.com/osniantonio/goexpert/12/1
