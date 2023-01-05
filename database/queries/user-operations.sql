-- name: CreateUser :one
insert into users
    (username, hashedPassword, name, email)
    values (@username, @hashedPassword, @name, @email)
        returning *;

-- name: GetUserDetails :one
select * from users
    where username= @username
        limit 1;

-- name: UpdateUserDetails :one
update users
    set
        hashedPassword= COALESCE(@hashedPassword, hashedPassword),
        name= COALESCE(@name, name),
        email= COALESCE(@email, email)
    where username= @username
        returning *;