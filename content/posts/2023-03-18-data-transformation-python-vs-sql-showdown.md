+++
title = "Data transformation - Python vs SQL showdown"
date = "2023-03-18"
path = "/posts/2023/03/data-transformation-python-vs-sql-showdown"

[taxonomies]
categories = [ "data-engineering",]
tags = []

+++

For most people, using SQL to transform data is a no-brainer, seeing it's a very versatile language, and doesn't have quite a steep learning curve compared to python. There are some cases where SQL is more suitable for a task, but the reverse can also happen as well.

For instance, given a string conversion problem:

- if a string occurs only one time, replace it with `#`
- if a string occurs multiple times, replace it with `&`

```
> one
###

> three
###&&

> Heartbreak hotel
&&&&&#&&&##&#&&#
```

A solution in python would be:

```python
from collections import Counter

s = "three"
s_counter = Counter(s)

output_str = ""
for char in s:
    if s_counter[char] > 1:
        output_str += "&"
    elif s_counter[char] == 1:
        output_str += "#"
```

But a solution in SQL is...(thanks Emily @data-engineering-discord!):

```sql
CREATE TABLE data (string_value TEXT);

INSERT INTO
  data
VALUES
  ('one'),
  ('three'),
  ('Heartbreak hotel');

select
  string_value,
  translate(
    lower(string_value),
    string_agg (
      chr,
      ''
      order by
        chr
    ),
    string_agg (
      subs,
      ''
      order by
        chr
    )
  )
from
  (
    select
      string_value,
      chr,
      case
        when count = 1 then '#'
        else '&'
      end as subs
    from
      (
        select
          p.*,
          count(*)
        from
          (
            select
              string_value,
              regexp_split_to_table (lower(string_value), '') as chr
            from
              data
          ) as p
        group by
          1,
          2
      ) as q
  ) as r
group by
  1
```

😱😱😱😱😱😱😱😱😱
