# MySQL
## tsv import
```
mysqlimport -u ユーザー名 --password=パスワード --local データベース名 テーブル名.txt
http://qiita.com/aosho235/items/f52f068d3634b0521cca

下記コマンドもあるみたい
LOAD DATA LOCAL INFILE 'ファイルパス' INTO TABLE user FIELDS TERMINATED BY '\t' LINES TERMINATED BY '\r\n';
http://www.infoscoop.org/blogjp/2014/07/23/about-load-data-infile/
```
