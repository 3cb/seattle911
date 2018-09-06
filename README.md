# seattle911

> Heatmap and Street View maps for Seattle, WA 911 calls

> Fire API Docs: https://dev.socrata.com/foundry/data.seattle.gov/grwu-wqtk

*** The following API has been deprecated - data stops after May 8, 2018
*** New dataset (https://data.seattle.gov/resource/xurz-654a.json) does not include latitude and longitude
> Police API Docs: https://dev.socrata.com/foundry/data.seattle.gov/pu5n-trf4

## Build Setup
for static files:

```
# move to static folder
cd /path/to/project/static

# install dependencies
npm install

# serve with hot reload at localhost:8080
npm run dev

# build for production with minification
npm run build
```

For detailed explanation on how things work, consult the [docs for vue-loader](http://vuejs.github.io/vue-loader).

for go files:

```
# from project root
go get
go build
./{binary_name}
```