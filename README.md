# RedCoins

RedCoins is an API to buy and sell bitcoins. It is possible to do that using dollars or reais. If you are an user of the time ADMIN you can also see reports of the transactions by user or day.

## Running

To run the RedCoins application first you need to create a mysql database and create the "users" and "reports" table

```
create table users(
id INTEGER NOT NULL AUTO_INCREMENT,
name VARCHAR(100) UNIQUE NOT NULL,
email VARCHAR(255) UNIQUE NOT NULL,
password VARCHAR(500) NOT NULL,
birth_date DATE NULL,
type VARCHAR(10) NOT NULL,
PRIMARY KEY(id));


create table reports (
id INTEGER NOT NULL AUTO_INCREMENT,
transaction_date DATE NULL,
bitcoins DOUBLE(20, 10) NOT NULL,
transaction VARCHAR(30) NOT NULL,
value DOUBLE(20, 10) NOT NULL,
currency VARCHAR(5) NOT NULL,
user_id INTEGER NOT NULL,
PRIMARY KEY(id),
FOREIGN KEY(user_id) REFERENCES users(id));
```

You will then have to change the [file](https://github.com/iaronaraujo/RedCoins/blob/master/lib/db.go) responsible for setting the database up, putting the correct values in the config fields.

After that, you can run the app.go file normally and it should open the application on localhost:3000

## Details

Even though you pass the date of operation when you buy or sell bitcoin, the quote is calculated based on the *current* date. This operation was done this way because it was not possible to get the bitcoin quote for other dates using the free API.

Check the wiki page for more details.
