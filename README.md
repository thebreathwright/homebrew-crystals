# Crystals Homebrew Tap

This tap publishes the `crystals` Homebrew formula and package assets.

- Formula: [`Formula/crystals.rb`](Formula/crystals.rb)
- Package files:
  - [`docs/crystals_package.md`](docs/crystals_package.md)
  - [`api/types.go`](api/types.go)
  - [`server/routes.go`](server/routes.go)
  - [`server/grax_consequence_contract.go`](server/grax_consequence_contract.go)
- Install with explicit Homebrew 6.0.0 tap trust as needed

## Install

```shell
brew tap thebreathwright/crystals
brew trust --formula thebreathwright/crystals/crystals
brew install crystals
```
