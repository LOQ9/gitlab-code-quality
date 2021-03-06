package model

var eslintCategory = map[string]string{
	"accessor-pairs":                "Bug Risk",
	"array-bracket-spacing":         "Clarity",
	"array-callback-return":         "Clarity",
	"arrow-body-style":              "Clarity",
	"arrow-parens":                  "Clarity",
	"arrow-spacing":                 "Clarity",
	"block-scoped-var":              "Bug Risk",
	"block-spacing":                 "Clarity",
	"brace-style":                   "Clarity",
	"callback-return":               "Bug Risk",
	"camelcase":                     "Clarity",
	"capitalized-comments":          "Clarity",
	"class-methods-use-this":        "Bug Risk",
	"comma-dangle":                  "Bug Risk",
	"comma-spacing":                 "Clarity",
	"comma-style":                   "Clarity",
	"complexity":                    "Complexity",
	"computed-property-spacing":     "Clarity",
	"consistent-this":               "Bug Risk",
	"consistent-return":             "Bug Risk",
	"constructor-super":             "Bug Risk",
	"curly":                         "Clarity",
	"default-case":                  "Bug Risk",
	"dot-location":                  "Clarity",
	"dot-notation":                  "Clarity",
	"eol-last":                      "Clarity",
	"eqeqeq":                        "Bug Risk",
	"func-call-spacing":             "Clarity",
	"func-name-matching":            "Clarity",
	"func-names":                    "Clarity",
	"func-style":                    "Clarity",
	"generator-star-spacing":        "Clarity",
	"global-require":                "Clarity",
	"guard-for-in":                  "Bug Risk",
	"handle-callback-err":           "Bug Risk",
	"id-blacklist":                  "Clarity",
	"id-length":                     "Clarity",
	"id-match":                      "Clarity",
	"indent":                        "Clarity",
	"init-declarations":             "Clarity",
	"jsx-quotes":                    "Clarity",
	"key-spacing":                   "Clarity",
	"keyword-spacing":               "Clarity",
	"line-comment-style":            "Clarity",
	"linebreak-style":               "Clarity",
	"lines-around-comment":          "Clarity",
	"lines-around-directive":        "Clarity",
	"max-depth":                     "Complexity",
	"max-len":                       "Complexity",
	"max-lines":                     "Complexity",
	"max-nested-callbacks":          "Complexity",
	"max-params":                    "Complexity",
	"max-statements":                "Complexity",
	"max-statements-per-line":       "Complexity",
	"multiline-ternary":             "Clarity",
	"new-cap":                       "Clarity",
	"new-parens":                    "Clarity",
	"newline-after-var":             "Clarity",
	"newline-before-return":         "Clarity",
	"newline-per-chained-call":      "Clarity",
	"no-alert":                      "Bug Risk",
	"no-array-constructor":          "Clarity",
	"no-await-in-loop":              "Bug Risk",
	"no-bitwise":                    "Clarity",
	"no-caller":                     "Compatibility",
	"no-case-declarations":          "Bug Risk",
	"no-catch-shadow":               "Bug Risk",
	"no-class-assign":               "Bug Risk",
	"no-cond-assign":                "Bug Risk",
	"no-confusing-arrow":            "Clarity",
	"no-console":                    "Bug Risk",
	"no-constant-condition":         "Bug Risk",
	"no-const-assign":               "Bug Risk",
	"no-continue":                   "Clarity",
	"no-control-regex":              "Bug Risk",
	"no-debugger":                   "Bug Risk",
	"no-delete-var":                 "Bug Risk",
	"no-div-regex":                  "Bug Risk",
	"no-dupe-args":                  "Bug Risk",
	"no-dupe-keys":                  "Bug Risk",
	"no-dupe-class-members":         "Bug Risk",
	"no-duplicate-case":             "Bug Risk",
	"no-duplicate-imports":          "Clarity",
	"no-else-return":                "Clarity",
	"no-empty":                      "Bug Risk",
	"no-empty-character-class":      "Bug Risk",
	"no-empty-function":             "Bug Risk",
	"no-empty-pattern":              "Bug Risk",
	"no-eq-null":                    "Bug Risk",
	"no-eval":                       "Security",
	"no-ex-assign":                  "Bug Risk",
	"no-extend-native":              "Bug Risk",
	"no-extra-bind":                 "Bug Risk",
	"no-extra-boolean-cast":         "Bug Risk",
	"no-extra-label":                "Bug Risk",
	"no-extra-parens":               "Bug Risk",
	"no-extra-semi":                 "Bug Risk",
	"no-fallthrough":                "Bug Risk",
	"no-floating-decimal":           "Clarity",
	"no-func-assign":                "Bug Risk",
	"no-global-assign":              "Bug Risk",
	"no-implicit-coercion":          "Bug Risk",
	"no-implicit-globals":           "Clarity",
	"no-implied-eval":               "Security",
	"no-inline-comments":            "Clarity",
	"no-inner-declarations":         "Compatibility",
	"no-invalid-regexp":             "Bug Risk",
	"no-invalid-this":               "Bug Risk",
	"no-irregular-whitespace":       "Compatibility",
	"no-iterator":                   "Compatibility",
	"no-label-var":                  "Bug Risk",
	"no-labels":                     "Bug Risk",
	"no-lone-blocks":                "Bug Risk",
	"no-lonely-if":                  "Bug Risk",
	"no-loop-func":                  "Bug Risk",
	"no-magic-numbers":              "Clarity",
	"no-mixed-operators":            "Clarity",
	"no-mixed-requires":             "Clarity",
	"no-mixed-spaces-and-tabs":      "Clarity",
	"no-multi-assign":               "Clarity",
	"no-multi-spaces":               "Bug Risk",
	"no-multi-str":                  "Compatibility",
	"no-multiple-empty-lines":       "Clarity",
	"no-negated-condition":          "Clarity",
	"no-nested-ternary":             "Bug Risk",
	"no-new":                        "Bug Risk",
	"no-new-func":                   "Clarity",
	"no-new-object":                 "Clarity",
	"no-new-require":                "Clarity",
	"no-new-wrappers":               "Bug Risk",
	"no-obj-calls":                  "Bug Risk",
	"no-octal":                      "Compatibility",
	"no-octal-escape":               "Compatibility",
	"no-param-reassign":             "Bug Risk",
	"no-path-concat":                "Bug Risk",
	"no-plusplus":                   "Bug Risk",
	"no-process-env":                "Bug Risk",
	"no-process-exit":               "Bug Risk",
	"no-proto":                      "Compatibility",
	"no-prototype-builtins":         "Compatibility",
	"no-redeclare":                  "Bug Risk",
	"no-regex-spaces":               "Bug Risk",
	"no-restricted-globals":         "Clarity",
	"no-restricted-imports":         "Clarity",
	"no-restricted-modules":         "Security",
	"no-restricted-properties":      "Bug Risk",
	"no-restricted-syntax":          "Bug Risk",
	"no-return-assign":              "Bug Risk",
	"no-return-await":               "Bug Risk",
	"no-script-url":                 "Security",
	"no-self-assign":                "Bug Risk",
	"no-self-compare":               "Bug Risk",
	"no-sequences":                  "Bug Risk",
	"no-shadow":                     "Bug Risk",
	"no-shadow-restricted-names":    "Bug Risk",
	"no-sparse-arrays":              "Bug Risk",
	"no-sync":                       "Bug Risk",
	"no-tabs":                       "Clarity",
	"no-template-curly-in-string":   "Clarity",
	"no-ternary":                    "Clarity",
	"no-this-before-super":          "Clarity",
	"no-throw-literal":              "Clarity",
	"no-trailing-spaces":            "Clarity",
	"no-undef":                      "Bug Risk",
	"no-undef-init":                 "Bug Risk",
	"no-undefined":                  "Compatibility",
	"no-underscore-dangle":          "Clarity",
	"no-unexpected-multiline":       "Bug Risk",
	"no-unmodified-loop-condition":  "Bug Risk",
	"no-unneeded-ternary":           "Clarity",
	"no-unreachable":                "Bug Risk",
	"no-unsafe-finally":             "Bug Risk",
	"no-unsafe-negation":            "Bug Risk",
	"no-unused-expressions":         "Bug Risk",
	"no-unused-vars":                "Bug Risk",
	"no-unused-labels":              "Clarity",
	"no-use-before-define":          "Compatibility",
	"no-useless-call":               "Bug Risk",
	"no-useless-computed-key":       "Clarity",
	"no-useless-concat":             "Bug Risk",
	"no-useless-constructor":        "Clarity",
	"no-useless-escape":             "Clarity",
	"no-useless-rename":             "Clarity",
	"no-useless-return":             "Clarity",
	"no-var":                        "Clarity",
	"no-void":                       "Compatibility",
	"no-warning-comments":           "Bug Risk",
	"no-whitespace-before-property": "Clarity",
	"no-with":                       "Compatibility",
	"object-curly-newline":          "Clarity",
	"object-curly-spacing":          "Clarity",
	"object-property-newline":       "Clarity",
	"object-shorthand":              "Clarity",
	"one-var":                       "Clarity",
	"one-var-declaration-per-line":  "Clarity",
	"operator-assignment":           "Clarity",
	"operator-linebreak":            "Clarity",
	"padded-blocks":                 "Clarity",
	"prefer-arrow-callback":         "Clarity",
	"prefer-const":                  "Clarity",
	"prefer-destructuring":          "Clarity",
	"prefer-numeric-literals":       "Clarity",
	"prefer-promise-reject-errors":  "Clarity",
	"prefer-rest-params":            "Clarity",
	"prefer-spread":                 "Clarity",
	"prefer-template":               "Clarity",
	"quote-props":                   "Clarity",
	"quotes":                        "Clarity",
	"radix":                         "Bug Risk",
	"require-await":                 "Clarity",
	"require-jsdoc":                 "Clarity",
	"require-yield":                 "Clarity",
	"rest-spread-spacing":           "Clarity",
	"semi":                          "Clarity",
	"semi-spacing":                  "Clarity",
	"sort-imports":                  "Clarity",
	"sort-vars":                     "Clarity",
	"space-before-blocks":           "Clarity",
	"space-before-function-paren":   "Clarity",
	"space-in-parens":               "Clarity",
	"space-infix-ops":               "Clarity",
	"space-unary-ops":               "Clarity",
	"spaced-comment":                "Clarity",
	"strict":                        "Clarity",
	"symbol-description":            "Clarity",
	"template-curly-spacing":        "Clarity",
	"template-tag-spacing":          "Clarity",
	"unicode-bom":                   "Clarity",
	"use-isnan":                     "Bug Risk",
	"valid-jsdoc":                   "Clarity",
	"valid-typeof":                  "Bug Risk",
	"vars-on-top":                   "Clarity",
	"wrap-iife":                     "Clarity",
	"wrap-regex":                    "Clarity",
	"yoda":                          "Clarity",

	"jest/no-disabled-tests":  "Bug Risk",
	"jest/no-focused-tests":   "Clarity",
	"jest/no-identical-title": "Clarity",
	"jest/valid-expect":       "Bug Risk",

	"jsdoc/check-param-names":                       "Clarity",
	"jsdoc/check-tag-names":                         "Clarity",
	"jsdoc/check-types":                             "Clarity",
	"jsdoc/newline-after-description":               "Clarity",
	"jsdoc/require-description-complete-sentence":   "Clarity",
	"jsdoc/require-hyphen-before-param-description": "Clarity",
	"jsdoc/require-param":                           "Clarity",
	"jsdoc/require-param-description":               "Clarity",
	"jsdoc/require-param-type":                      "Clarity",
	"jsdoc/require-returns-description":             "Clarity",
	"jsdoc/require-returns-type":                    "Clarity",

	"lodash/callback-binding":        "Compatibility",
	"lodash/chain-style":             "Clarity",
	"lodash/chaining":                "Clarity",
	"lodash/collection-method-value": "Bug Risk",
	"lodash/collection-return":       "Bug Risk",
	"lodash/consistent-compose":      "Clarity",
	"lodash/identity-shorthand":      "Clarity",
	"lodash/import-scope":            "Clarity",
	"lodash/matches-prop-shorthand":  "Clarity",
	"lodash/no-commit":               "Bug Risk",
	"lodash/no-double-unwrap":        "Bug Risk",
	"lodash/no-extra-args":           "Clarity",
	"lodash/no-unbound-this":         "Bug Risk",
	"lodash/path-style":              "Clarity",
	"lodash/prefer-compact":          "Bug Risk",
	"lodash/prefer-constant":         "Bug Risk",
	"lodash/prefer-filter":           "Bug Risk",
	"lodash/prefer-flat-map":         "Bug Risk",
	"lodash/prefer-get":              "Bug Risk",
	"lodash/prefer-includes":         "Clarity",
	"lodash/prefer-invoke-map":       "Bug Risk",
	"lodash/prefer-is-nil":           "Bug Risk",
	"lodash/prefer-lodash-chain":     "Clarity",
	"lodash/prefer-lodash-method":    "Clarity",
	"lodash/prefer-lodash-typecheck": "Bug Risk",
	"lodash/prefer-map":              "Bug Risk",
	"lodash/prefer-matches":          "Clarity",
	"lodash/prefer-noop":             "Clarity",
	"lodash/prefer-over-quantifer":   "Bug Risk",
	"lodash/prefer-reject":           "Bug Risk",
	"lodash/prefer-startswith":       "Bug Risk",
	"lodash/prefer-thru":             "Bug Risk",
	"lodash/prefer-times":            "Clarity",
	"lodash/prefer-wrapper-method":   "Bug Risk",
	"lodash/preferred-alias":         "Clarity",
	"lodash/prop-shorthand":          "Clarity",
	"lodash/unwrap":                  "Bug Risk",

	"mocha/handle-done-callback":     "Clarity",
	"mocha/max-top-level-suites":     "Clarity",
	"mocha/no-exclusive-tests":       "Bug Risk",
	"mocha/no-global-tests":          "Clarity",
	"mocha/no-hooks":                 "Clarity",
	"mocha/no-hooks-for-single-case": "Clarity",
	"mocha/no-identical-title":       "Clarity",
	"mocha/no-mocha-arrows":          "Bug Risk",
	"mocha/no-nested-suites":         "Clarity",
	"mocha/no-pending-tests":         "Bug Risk",
	"mocha/no-return-and-callback":   "Clarity",
	"mocha/no-sibling-hooks":         "Clarity",
	"mocha/no-skipped-tests":         "Bug Risk",
	"mocha/no-synchronous-tests":     "Bug Risk",
	"mocha/no-top-level-hooks":       "Clarity",
	"mocha/valid-test-description":   "Clarity",
	"mocha/valid-suite-description":  "Clarity",

	"mongodb/check-addtoset-updates":     "Bug Risk",
	"mongodb/check-current-date-updates": "Bug Risk",
	"mongodb/check-deprecated-calls":     "Bug Risk",
	"mongodb/check-deprecated-updates":   "Bug Risk",
	"mongodb/check-insert-calls":         "Bug Risk",
	"mongodb/check-minmax-updates":       "Bug Risk",
	"mongodb/check-numeric-updates":      "Bug Risk",
	"mongodb/check-pop-updates":          "Bug Risk",
	"mongodb/check-pull-updates":         "Bug Risk",
	"mongodb/check-push-updates":         "Bug Risk",
	"mongodb/check-query-calls":          "Bug Risk",
	"mongodb/check-remove-calls":         "Bug Risk",
	"mongodb/check-rename-updates":       "Bug Risk",
	"mongodb/check-set-updates":          "Bug Risk",
	"mongodb/check-unset-updates":        "Bug Risk",
	"mongodb/check-update-calls":         "Bug Risk",
	"mongodb/no-replace":                 "Bug Risk",

	"react-native/no-unused-styles":          "Clarity",
	"react-native/split-platform-components": "Clarity",
	"react-native/no-inline-styles":          "Clarity",
	"react-native/no-color-literals":         "Clarity",

	"security/detect-buffer-noassert":                "Security",
	"security/detect-child-process":                  "Security",
	"security/detect-disable-mustache-escape":        "Security",
	"security/detect-eval-with-expression":           "Security",
	"security/detect-new-buffer":                     "Security",
	"security/detect-no-csrf-before-method-override": "Security",
	"security/detect-non-literal-fs-filename":        "Security",
	"security/detect-non-literal-regexp":             "Security",
	"security/detect-non-literal-require":            "Security",
	"security/detect-object-injection":               "Security",
	"security/detect-possible-timing-attacks":        "Security",
	"security/detect-pseudoRandomBytes":              "Security",
	"security/detect-unsafe-regex":                   "Security",
}
