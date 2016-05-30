# RedashToEmail

This should run as an automated build (circleci, but other should work as well) it logs in to your redash server via a url, and grabs a screenshot of a specificed html element (via an xpath) and uploads it to amazon S3.

## Example 

change your circle.yml to run 
```
- go run main.go -url http://<yourredash url/<your redash dashboard> -xpath "//*[@id="new_blob"]/div[2]/span[3]/a"
```

## License


MIT (see [LICENSE](https://github.com/streamrail/concurrent-map/blob/master/LICENSE.txt) file)
