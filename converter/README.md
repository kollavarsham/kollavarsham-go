# [kollavarsham](http://kollavarsham.org/)

[![Circle CI Status](https://img.shields.io/circleci/build/github/kollavarsham/kollavarsham-js?label=CircleCI)](https://app.circleci.com/pipelines/github/kollavarsham/kollavarsham-js) [![Coverage Status](https://img.shields.io/coveralls/github/kollavarsham/kollavarsham-js?label=Coveralls)](https://coveralls.io/github/kollavarsham/kollavarsham-js?branch=main) [![GitHub Actions Status](https://github.com/kollavarsham/kollavarsham-js/actions/workflows/ci.yml/badge.svg)](https://github.com/kollavarsham/kollavarsham-js/actions/workflows/ci.yml?query=branch%3Amain) [![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fkollavarsham%2Fkollavarsham-js.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fkollavarsham%2Fkollavarsham-js?ref=badge_shield)

> Convert Gregorian date to Kollavarsham date and vice versa

## Install

### TypeScript/JavaScript/Node.js [![NPM version](https://badge.fury.io/js/kollavarsham.svg)](https://www.npmjs.com/package/kollavarsham)

```sh
$ npm install kollavarsham
```

### Python [![PyPI version](https://img.shields.io/pypi/v/kollavarsham)](https://pypi.org/project/kollavarsham/)

```sh
$ pip install kollavarsham
```

### Go [![Go project version](https://badge.fury.io/go/github.com%2Fkollavarsham%2Fkollavarsham-go%2Fconverter%2Fv2.svg)](https://pkg.go.dev/github.com/kollavarsham/kollavarsham-go/converter/v2)

```sh
go get github.com/kollavarsham/kollavarsham-go/converter
```

### Java [![Maven version](https://img.shields.io/maven-central/v/org.kollavarsham.converter/kollavarsham-converter)](https://search.maven.org/artifact/org.kollavarsham.converter/kollavarsham-converter)

```xml
<dependency>
  <groupId>org.kollavarsham.converter</groupId>
  <artifactId>kollavarsham-converter</artifactId>
  <version>2.0.1</version>
</dependency>
```

### C#/dotnet [![NuGet version](https://badge.fury.io/nu/KollavarshamOrg.Converter.svg)](https://www.nuget.org/packages/KollavarshamOrg.Converter)

```sh
$ dotnet add package KollavarshamOrg.Converter
```

## Usage

Refer [the samples repository](https://github.com/kollavarsham/kollavarsham-samples) for working examples.

### TypeScript/JavaScript/Node.js

```js
import { Kollavarsham } from 'kollavarsham';

const options = {
  system: 'SuryaSiddhanta',
  latitude: 10,
  longitude: 76.2
};

const kollavarsham = new Kollavarsham(options);

const today = kollavarsham.fromGregorianDate(new Date());

console.log(today.year, today.mlMasaName, today.date, `(${today.mlNaksatraName})`);
```

### Python

```python
import datetime
import pytz
import kollavarsham

now = pytz.utc.localize(datetime.datetime.utcnow())
kv = kollavarsham.Kollavarsham(latitude=10, longitude=76.2, system="SuryaSiddhanta")

today = kv.from_gregorian_date(date=now)
print(today.year, today.ml_masa_name, today.date, '(' + today.naksatra.ml_malayalam + ')')
```

### Go

```go
package main

import (
	"fmt"
	"time"

	"github.com/kollavarsham/kollavarsham-go/converter/v2"
)

func main() {

	latitude := float64(23.2)
	longitude := float64(75.8)
	system := "SuryaSiddhanta"
	kv := converter.NewKollavarsham(&converter.Settings{
		Latitude:  &latitude,
		Longitude: &longitude,
		System:    &system,
	})

	now := time.Now()
	today := kv.FromGregorianDate(&now)

	fmt.Printf("Today in Malayalam Year: %v %v %v (%v)\n", *today.Year(), *today.MlMasaName(), *today.Date(), *today.MlNaksatraName())
}
```

### Java

```java
package org.kollavarsham.tester;

import java.time.Instant;

import org.kollavarsham.converter.Kollavarsham;
import org.kollavarsham.converter.KollavarshamDate;
import org.kollavarsham.converter.Settings;
import org.kollavarsham.converter.Settings.Builder;

public class App {
    public static void main( final String[] args) {
        final Settings settings = new Builder().latitude(10).longitude(76.2).system("SuryaSiddhanta").build();
        final Kollavarsham kv = new Kollavarsham(settings);
        final KollavarshamDate today = kv.fromGregorianDate(Instant.now());
        System.out.println( today.getYear() + today.getMlMasaName() + today.getDate() + '(' + today.getMlNaksatraName() + ')' );
    }
}
```

### C#/dotnet

```csharp
using System;

namespace KollavarshamOrg.Tester
{
    class Program
    {
        static void Main(string[] args)
        {
            var settings = new Settings {
                Latitude = 10,
                Longitude = 76.2,
                System = "SuryaSiddhanta"
            };
            var kv = new Kollavarsham(settings);
            var today = kv.FromGregorianDate(DateTime.Now);
            Console.WriteLine($"{today.Year.ToString()} {today.MlMasaName} {today.Date.ToString()} ({today.MlNaksatraName})");
        }
    }
}
```

## Documentation

### TypeScript/JavaScript/Node.js

Check out the [Kollavarsham class](https://kollavarsham.org/kollavarsham-js/module-kollavarsham.Kollavarsham.html) within the API documentation as this is the entry point into the library.

## Release History

Check out the history at [GitHub Releases](https://github.com/kollavarsham/kollavarsham-js/releases)

## License

Copyright (c) 2014-2023 The Kollavarsham Team. Licensed under the [MIT license](http://kollavarsham.org/LICENSE.txt).

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fkollavarsham%2Fkollavarsham-js.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fkollavarsham%2Fkollavarsham-js?ref=badge_large)
