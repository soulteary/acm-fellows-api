# ACM Fellows API

API of ACM Fellows.

## Quick Start

Run go-nameparser service first:

```bash
docker run --rm -it -p 8080:8080 -p 8081:8081 soulteary/go-nameparser
```

Run the following command to fetch and update the ACM csv file:

```bash
go run .
```

## Credits

- [emeryberger/CSrankings/util/acm-fellow-scraper.py](https://github.com/emeryberger/CSrankings/blob/gh-pages/util/acm-fellow-scraper.py) origin implementation.
- [soulteary/go-nameparser](https://github.com/soulteary/go-nameparser) go name parser service.
- [derek73/python-nameparser](https://github.com/derek73/python-nameparser) python name parser utils.
