/*
 * Cadence - The resource-oriented smart contract programming language
 *
 * Copyright Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package checker

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/onflow/cadence/runtime/sema"
	"github.com/onflow/cadence/runtime/stdlib"
)

func TestCheckTypeArguments(t *testing.T) {

	t.Parallel()

	t.Run("capability, no instantiation", func(t *testing.T) {

		t.Parallel()

		checker, err := parseAndCheckWithTestValue(t,
			`
              let cap: Capability = test
            `,
			&sema.CapabilityType{},
		)

		require.NoError(t, err)

		capType := RequireGlobalValue(t, checker.Elaboration, "cap")

		require.IsType(t,
			&sema.CapabilityType{},
			capType,
		)
		require.Nil(t, capType.(*sema.CapabilityType).BorrowType)
	})

	t.Run("inclusive range, instantiation with more than arguments", func(t *testing.T) {

		t.Parallel()

		baseValueActivation := sema.NewVariableActivation(sema.BaseValueActivation)
		baseValueActivation.DeclareValue(stdlib.InclusiveRangeConstructorFunction)

		_, err := ParseAndCheckWithOptions(t,
			`
              let inclusiveRange: InclusiveRange<Int, String> = InclusiveRange(1, 10)
            `,
			ParseAndCheckOptions{
				Config: &sema.Config{
					BaseValueActivation: baseValueActivation,
				},
			},
		)

		errs := RequireCheckerErrors(t, err, 2)

		assert.IsType(t, &sema.InvalidTypeArgumentCountError{}, errs[0])
		assert.IsType(t, &sema.MissingTypeArgumentError{}, errs[1])
	})

	t.Run("capability, instantiation with no arguments", func(t *testing.T) {

		t.Parallel()

		_, err := parseAndCheckWithTestValue(t,
			`
              let cap: Capability<> = test
            `,
			&sema.CapabilityType{},
		)

		errs := RequireCheckerErrors(t, err, 1)

		assert.IsType(t, &sema.InvalidTypeArgumentCountError{}, errs[0])
	})

	t.Run("capability, instantiation with one argument, reference", func(t *testing.T) {

		t.Parallel()

		checker, err := parseAndCheckWithTestValue(t,
			`
              let cap: Capability<&Int> = test
            `,
			&sema.CapabilityType{
				BorrowType: &sema.ReferenceType{
					Type: sema.IntType,
				},
			},
		)
		require.NoError(t, err)

		assert.Equal(t,
			&sema.CapabilityType{
				BorrowType: &sema.ReferenceType{
					Type: sema.IntType,
				},
			},
			RequireGlobalValue(t, checker.Elaboration, "cap"),
		)
	})

	t.Run("capability, instantiation with one argument, non-reference", func(t *testing.T) {

		t.Parallel()

		_, err := parseAndCheckWithTestValue(t,
			`
              let cap: Capability<Int> = test
            `,
			&sema.CapabilityType{
				BorrowType: sema.IntType,
			},
		)

		errs := RequireCheckerErrors(t, err, 1)

		assert.IsType(t, &sema.TypeMismatchError{}, errs[0])
	})

	t.Run("capability, instantiation with two arguments", func(t *testing.T) {

		t.Parallel()

		_, err := parseAndCheckWithTestValue(t,
			`
              let cap: Capability<&Int, &String> = test
            `,
			&sema.CapabilityType{},
		)

		errs := RequireCheckerErrors(t, err, 1)

		assert.IsType(t, &sema.InvalidTypeArgumentCountError{}, errs[0])
	})
}

type checkParameterizedTypeIsInstantiatedTestCase struct {
	name string
	code string
}

func TestCheckParameterizedTypeIsInstantiated(t *testing.T) {

	t.Parallel()

	baseValueActivation := sema.NewVariableActivation(sema.BaseValueActivation)
	baseValueActivation.DeclareValue(stdlib.InclusiveRangeConstructorFunction)
	options := ParseAndCheckOptions{
		Config: &sema.Config{
			BaseValueActivation: baseValueActivation,
		},
	}

	runTestCase := func(t *testing.T, testCase checkParameterizedTypeIsInstantiatedTestCase) {
		t.Run(testCase.name, func(t *testing.T) {

			t.Parallel()

			_, err := ParseAndCheckWithOptions(t,
				testCase.code,
				options,
			)

			errs := RequireCheckerErrors(t, err, 1)

			assert.IsType(t, &sema.MissingTypeArgumentError{}, errs[0])
		})
	}

	testCases := []checkParameterizedTypeIsInstantiatedTestCase{
		{
			name: "InclusiveRange",
			code: "let inclusiveRange: InclusiveRange = InclusiveRange(1, 10)",
		},
		{
			name: "VariableSizedArray with InclusiveRange",
			code: "let r: [InclusiveRange] = []",
		},
		{
			name: "ConstantSizedType with InclusiveRange",
			code: "let r: [InclusiveRange; 2] = [InclusiveRange(1, 2), InclusiveRange(3, 4)]",
		},
		{
			name: "OptionalType with InclusiveRange",
			code: "let r: InclusiveRange? = nil",
		},
		{
			name: "DictionaryType with InclusiveRange",
			code: "let r: {Int: InclusiveRange} = {}",
		},
		{
			name: "Struct with InclusiveRange",
			code: `
				pub struct Foo {
					pub let a: InclusiveRange

					init() {
						self.a = InclusiveRange(1, 10)
					}
				}
			`,
		},
	}

	for _, test := range testCases {
		runTestCase(t, test)
	}
}

func TestCheckTypeArgumentSubtyping(t *testing.T) {

	t.Parallel()

	t.Run("Capability<&Int> is a subtype of Capability", func(t *testing.T) {

		t.Parallel()

		checker, err := parseAndCheckWithTestValue(t,
			`
              let cap: Capability<&Int> = test
              let cap2: Capability = cap
            `,
			&sema.CapabilityType{
				BorrowType: &sema.ReferenceType{
					Type: sema.IntType,
				},
			},
		)
		require.NoError(t, err)

		capType := RequireGlobalValue(t, checker.Elaboration, "cap")
		require.IsType(t,
			&sema.CapabilityType{},
			capType,
		)
		require.Equal(t,
			&sema.ReferenceType{
				Type: sema.IntType,
			},
			capType.(*sema.CapabilityType).BorrowType,
		)

		cap2Type := RequireGlobalValue(t, checker.Elaboration, "cap2")
		require.IsType(t,
			&sema.CapabilityType{},
			cap2Type,
		)
		require.Nil(t,
			cap2Type.(*sema.CapabilityType).BorrowType,
		)
	})

	t.Run("Capability<&Int> is a subtype of Capability<&Int>", func(t *testing.T) {

		t.Parallel()

		checker, err := parseAndCheckWithTestValue(t,
			`
              let cap: Capability<&Int> = test
              let cap2: Capability<&Int> = cap
            `,
			&sema.CapabilityType{
				BorrowType: &sema.ReferenceType{
					Type: sema.IntType,
				},
			},
		)
		require.NoError(t, err)

		assert.Equal(t,
			&sema.CapabilityType{
				BorrowType: &sema.ReferenceType{
					Type: sema.IntType,
				},
			},
			RequireGlobalValue(t, checker.Elaboration, "cap"),
		)

		assert.Equal(t,
			&sema.CapabilityType{
				BorrowType: &sema.ReferenceType{
					Type: sema.IntType,
				},
			},
			RequireGlobalValue(t, checker.Elaboration, "cap2"),
		)
	})

	t.Run("Capability is not a subtype of Capability<&Int>", func(t *testing.T) {

		t.Parallel()

		checker, err := parseAndCheckWithTestValue(t,
			`
              let cap: Capability = test
              let cap2: Capability<&Int> = cap
            `,
			&sema.CapabilityType{},
		)
		require.NotNil(t, checker)

		capType := RequireGlobalValue(t, checker.Elaboration, "cap")
		require.IsType(t,
			&sema.CapabilityType{},
			capType,
		)
		require.Nil(t, capType.(*sema.CapabilityType).BorrowType)

		cap2Type := RequireGlobalValue(t, checker.Elaboration, "cap2")
		require.IsType(t,
			&sema.CapabilityType{},
			cap2Type,
		)
		require.Equal(t,
			&sema.ReferenceType{
				Type: sema.IntType,
			},
			cap2Type.(*sema.CapabilityType).BorrowType,
		)

		errs := RequireCheckerErrors(t, err, 1)

		assert.IsType(t, &sema.TypeMismatchError{}, errs[0])
	})

	t.Run("Capability<&String> is not a subtype of Capability<&Int>", func(t *testing.T) {

		t.Parallel()

		checker, err := parseAndCheckWithTestValue(t,
			`
              let cap: Capability<&String> = test
              let cap2: Capability<&Int> = cap
            `,
			&sema.CapabilityType{
				BorrowType: &sema.ReferenceType{
					Type: sema.StringType,
				},
			},
		)
		require.NotNil(t, checker)

		assert.Equal(t,
			&sema.CapabilityType{
				BorrowType: &sema.ReferenceType{
					Type: sema.StringType,
				},
			},
			RequireGlobalValue(t, checker.Elaboration, "cap"),
		)

		assert.Equal(t,
			&sema.CapabilityType{
				BorrowType: &sema.ReferenceType{
					Type: sema.IntType,
				},
			},
			RequireGlobalValue(t, checker.Elaboration, "cap2"),
		)

		errs := RequireCheckerErrors(t, err, 1)

		assert.IsType(t, &sema.TypeMismatchError{}, errs[0])
	})
}
