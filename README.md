jsonget
-------

`jsonget` is a command-line tool for extracting values out from JSON. This is
useful, for example, when you're `curl`ing a JSON api and just want to get a
single value from it:

```bash
% curl http://openweathermap.org/data/2.0/weather/city/524901 | jsonget main.temp
259.92
```

Or multiple newline-separated values:

```bash
% curl http://openweathermap.org/data/2.0/weather/city/524901 | jsonget name main.temp wind.speed
Moscow
259.92
2
```

Installation
============

    go get github.com/tysontate/jsonget

Usage
=====

Given `my.json`:

```json
{
  "foo": true,
  "bar": {
    "baz": 5
  },
  "neat": "You bet.",
  "stuff": {
    "things": ["cheese", "barley", "corn"]
  }
}
```

`jsonget` can read from files:

```bash
% jsonget --file my.json foo
true
```

Or from stdin:

```bash
% cat my.json | jsonget foo
true
```

JSON strings are returned without surrounding quotes:

```bash
% cat my.json | jsonget neat
You bet.
```

Arrays can be accessed:

```base
% cat my.json | jsonget stuff.things[2] 
corn
```

And JSON objects and arrays are returned as JSON:

```bash
% cat my.json | jsonget bar
{"baz":5}
```

To Do
=====
 - Add wildcard support for arrays, example:

```
> echo "[{product:'coffee', price:2.10}, {product:'beer', price:3.80}]" | jsonget '*.price'

[2.10, 3.80]
```
