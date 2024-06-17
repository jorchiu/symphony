<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# helm

```go
import "github.com/eclipse-symphony/symphony/packages/testutils/expectations/helm"
```

## Index

- [Variables](<#variables>)
- [type HelmExpectation](<#HelmExpectation>)
  - [func MustNew\(name, namespace string, opts ...Option\) \*HelmExpectation](<#MustNew>)
  - [func MustNewAbsent\(name, namespace string, opts ...Option\) \*HelmExpectation](<#MustNewAbsent>)
  - [func NewExpectation\(pattern, namespace string, opts ...Option\) \(\*HelmExpectation, error\)](<#NewExpectation>)
  - [func \(e \*HelmExpectation\) AsGomegaSubject\(\) func\(context.Context\) \(interface\{\}, error\)](<#HelmExpectation.AsGomegaSubject>)
  - [func \(he \*HelmExpectation\) Description\(\) string](<#HelmExpectation.Description>)
  - [func \(he \*HelmExpectation\) Id\(\) string](<#HelmExpectation.Id>)
  - [func \(e \*HelmExpectation\) ToGomegaMatcher\(\) gomega.GomegaMatcher](<#HelmExpectation.ToGomegaMatcher>)
  - [func \(he \*HelmExpectation\) Verify\(c context.Context\) error](<#HelmExpectation.Verify>)
- [type ListRunner](<#ListRunner>)
- [type Option](<#Option>)
  - [func WithDescription\(description string\) Option](<#WithDescription>)
  - [func WithListClientBuilder\(builder func\(\) \(ListRunner, error\)\) Option](<#WithListClientBuilder>)
  - [func WithLogger\(logger func\(format string, args ...interface\{\}\)\) Option](<#WithLogger>)
  - [func WithReleaseCondition\(condition types.Condition\) Option](<#WithReleaseCondition>)
  - [func WithReleaseListCondition\(condition types.Condition\) Option](<#WithReleaseListCondition>)
  - [func WithRemoved\(removed bool\) Option](<#WithRemoved>)
  - [func WithValueCondition\(condition types.Condition\) Option](<#WithValueCondition>)
  - [func WithValueListCondition\(condition types.Condition\) Option](<#WithValueListCondition>)


## Variables

<a name="DeployedCondition"></a>

```go
var (
    DeployedCondition = jq.Equality(".info.status", "deployed")
    FailedCondition   = jq.Equality(".info.status", "failed")
)
```

<a name="HelmExpectation"></a>
## type [HelmExpectation](<https://github.com/eclipse-symphony/symphony/blob/main/packages/testutils/expectations/helm/resource.go#L27-L53>)



```go
type HelmExpectation struct {
    // contains filtered or unexported fields
}
```

<a name="MustNew"></a>
### func [MustNew](<https://github.com/eclipse-symphony/symphony/blob/main/packages/testutils/expectations/helm/resource.go#L96>)

```go
func MustNew(name, namespace string, opts ...Option) *HelmExpectation
```

MustNew creates a new helm expectation. It panics if the expectation cannot be created.

<a name="MustNewAbsent"></a>
### func [MustNewAbsent](<https://github.com/eclipse-symphony/symphony/blob/main/packages/testutils/expectations/helm/resource.go#L105>)

```go
func MustNewAbsent(name, namespace string, opts ...Option) *HelmExpectation
```

NewPresent creates a new helm expectation that expects the release to be present.

<a name="NewExpectation"></a>
### func [NewExpectation](<https://github.com/eclipse-symphony/symphony/blob/main/packages/testutils/expectations/helm/resource.go#L66>)

```go
func NewExpectation(pattern, namespace string, opts ...Option) (*HelmExpectation, error)
```

NewExpectation creates a new helm expectation.

<a name="HelmExpectation.AsGomegaSubject"></a>
### func \(\*HelmExpectation\) [AsGomegaSubject](<https://github.com/eclipse-symphony/symphony/blob/main/packages/testutils/expectations/helm/gomega.go#L15>)

```go
func (e *HelmExpectation) AsGomegaSubject() func(context.Context) (interface{}, error)
```



<a name="HelmExpectation.Description"></a>
### func \(\*HelmExpectation\) [Description](<https://github.com/eclipse-symphony/symphony/blob/main/packages/testutils/expectations/helm/resource.go#L238>)

```go
func (he *HelmExpectation) Description() string
```



<a name="HelmExpectation.Id"></a>
### func \(\*HelmExpectation\) [Id](<https://github.com/eclipse-symphony/symphony/blob/main/packages/testutils/expectations/helm/resource.go#L160>)

```go
func (he *HelmExpectation) Id() string
```

Id implements types.Expectation.

<a name="HelmExpectation.ToGomegaMatcher"></a>
### func \(\*HelmExpectation\) [ToGomegaMatcher](<https://github.com/eclipse-symphony/symphony/blob/main/packages/testutils/expectations/helm/gomega.go#L21>)

```go
func (e *HelmExpectation) ToGomegaMatcher() gomega.GomegaMatcher
```



<a name="HelmExpectation.Verify"></a>
### func \(\*HelmExpectation\) [Verify](<https://github.com/eclipse-symphony/symphony/blob/main/packages/testutils/expectations/helm/resource.go#L135>)

```go
func (he *HelmExpectation) Verify(c context.Context) error
```

Verify implements types.Expectation.

<a name="ListRunner"></a>
## type [ListRunner](<https://github.com/eclipse-symphony/symphony/blob/main/packages/testutils/expectations/helm/resource.go#L24-L26>)



```go
type ListRunner interface {
    Run() ([]*release.Release, error)
}
```

<a name="Option"></a>
## type [Option](<https://github.com/eclipse-symphony/symphony/blob/main/packages/testutils/expectations/helm/resource.go#L55>)



```go
type Option func(*HelmExpectation)
```

<a name="WithDescription"></a>
### func [WithDescription](<https://github.com/eclipse-symphony/symphony/blob/main/packages/testutils/expectations/helm/options.go#L43>)

```go
func WithDescription(description string) Option
```



<a name="WithListClientBuilder"></a>
### func [WithListClientBuilder](<https://github.com/eclipse-symphony/symphony/blob/main/packages/testutils/expectations/helm/options.go#L12>)

```go
func WithListClientBuilder(builder func() (ListRunner, error)) Option
```



<a name="WithLogger"></a>
### func [WithLogger](<https://github.com/eclipse-symphony/symphony/blob/main/packages/testutils/expectations/helm/options.go#L49>)

```go
func WithLogger(logger func(format string, args ...interface{})) Option
```



<a name="WithReleaseCondition"></a>
### func [WithReleaseCondition](<https://github.com/eclipse-symphony/symphony/blob/main/packages/testutils/expectations/helm/options.go#L31>)

```go
func WithReleaseCondition(condition types.Condition) Option
```



<a name="WithReleaseListCondition"></a>
### func [WithReleaseListCondition](<https://github.com/eclipse-symphony/symphony/blob/main/packages/testutils/expectations/helm/options.go#L37>)

```go
func WithReleaseListCondition(condition types.Condition) Option
```



<a name="WithRemoved"></a>
### func [WithRemoved](<https://github.com/eclipse-symphony/symphony/blob/main/packages/testutils/expectations/helm/options.go#L6>)

```go
func WithRemoved(removed bool) Option
```

WithRemoved specifies whether the release is expected to be present or not.

<a name="WithValueCondition"></a>
### func [WithValueCondition](<https://github.com/eclipse-symphony/symphony/blob/main/packages/testutils/expectations/helm/options.go#L24>)

```go
func WithValueCondition(condition types.Condition) Option
```



<a name="WithValueListCondition"></a>
### func [WithValueListCondition](<https://github.com/eclipse-symphony/symphony/blob/main/packages/testutils/expectations/helm/options.go#L18>)

```go
func WithValueListCondition(condition types.Condition) Option
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)