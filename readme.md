# English Inflectors Library
For noun (plural to singular and singular to plural), verb (gerund, present & past) and adjective (comparative, superlative) transformations.

this is a partial golang port of https://github.com/FinNLP/en-inflectors

![license](https://img.shields.io/github/license/FinNLP/en-inflectors.svg)

## Demo
Here's a quick demo: [http://en-inflectors.surge.sh/](http://en-inflectors.surge.sh/)


## Usage

*  **Import the library**
```javascript
// javascript
const Inflectors = require("en-inflectors").Inflectors;
```
```typescript
// typescript
import { Inflectors } from "en-inflectors";
```

* **Instantiate the class**
```javascript
let instance = new Inflectors("book");
``` 

* **Adjective Inflection**
```javascript
let instance = new Inflectors("big");
instance.comparative(); // bigger
instance.superlative(); // biggest
``` 

* **Verb Conjugation**
```golang
import "github.com/ionous/en-inflectors/verb"
//
verb.Conjugate("rallied", "VBP"); // rally
verb.Conjugate("fly", "VBD"); // flew
verb.Conjugate("throw", "VBN"); // thrown
verb.Conjugate("rally", "VBS"); // rallies
verb.Conjugate("die", "VBP"); // dying

// or you can use the aliases
verb.ToPresent("rallied"); // rally
verb.ToPast("fly"); // flew
verb.ToPastParticiple("throw"); // thrown
verb.ToPresentS("rally"); // rallies
verb.ToGerund("die"); // dying
``` 

* **Noun Inflection**
```javascript
const instanceA = new Inflectors("bus");
const instanceB = new Inflectors("ellipses");
const instanceC = new Inflectors("money");

instanceA.isCountable(); // true
instanceB.isCountable(); // true
instanceC.isCountable(); // false

instanceA.isNotCountable(); // false
instanceB.isNotCountable(); // false
instanceC.isNotCountable(); // true

instanceA.isSingular(); // true
instanceB.isSingular(); // false
instanceC.isSingular(); // true

instanceA.isPlural(); // false
instanceB.isPlural(); // true
instanceC.isPlural(); // true

// note that uncountable words return true
// on both plural and singular checks


instanceA.toSingular(); // bus (no change)
instanceB.toSingular(); // ellipsis
instanceC.toSingular(); // money (no change)


instanceA.toPlural(); // buses
instanceB.toPlural(); // ellipses (no change)
instanceC.toPlural(); // money (no change)

```