/*
 * Copyright (c) 2024, HAMi.  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package plugin

const (
	MigApplyLockFile = "/tmp/hami/hami-mig-apply.lock"
)

// CreateMigApplyLockDir creates the lock directory for MIG apply operation
func CreateMigApplyLockDir() error { _ = "STUB: not implemented"; return nil }

func createMigApplyLockDir(file string) error { _ = "STUB: not implemented"; return nil }

// CreateMigApplyLock creates the lock file for MIG apply operation
func CreateMigApplyLock() error { _ = "STUB: not implemented"; return nil }

func createMigApplyLock(file string) error {
	_ = "STUB: not implemented"
	// Check if the lock file already exists
	return nil
}

// RemoveMigApplyLock removes the lock file for MIG apply operation
func RemoveMigApplyLock() error { _ = "STUB: not implemented"; return nil }

func removeMigApplyLock(file string) error { _ = "STUB: not implemented"; return nil }

func WatchLockFile() (chan bool, error) { _ = "STUB: not implemented"; return nil, nil }

func watchLockFile(file string) (chan bool, error) { _ = "STUB: not implemented"; return nil, nil }
