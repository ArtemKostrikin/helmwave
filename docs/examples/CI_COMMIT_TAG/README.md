# CI_COMMIT_TAG  [ [source] ](https://github.com/zhilyaev/helmwave/tree/main/docs/examples/CI_COMMIT_TAG)

Project Structure

```
.
├── README.md
├── helmwave.yml.tpl
└── values.yml

```



```yaml
{% include_relative helmwave.yml.tpl %}
```

```yaml
{% capture my_include %}{% include values.yml %}{% endcapture %}
{{ my_include | markdownify }}
```
