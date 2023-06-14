
access(all) struct AuthAccount {

    /// The address of the account.
    access(all) let address: Address

    /// The FLOW balance of the default vault of this account.
    access(all) let balance: UFix64

    /// The FLOW balance of the default vault of this account that is available to be moved.
    access(all) let availableBalance: UFix64

    /// The current amount of storage used by the account in bytes.
    access(all) let storageUsed: UInt64

    /// The storage capacity of the account in bytes.
    access(all) let storageCapacity: UInt64

    /// The contracts deployed to the account.
    access(all) let contracts: AuthAccount.Contracts

    /// The keys assigned to the account.
    access(all) let keys: AuthAccount.Keys

    /// The inbox allows bootstrapping (sending and receiving) capabilities.
    access(all) let inbox: AuthAccount.Inbox

    /// The capabilities of the account.
    access(all) let capabilities: AuthAccount.Capabilities

    /// All public paths of this account.
    access(all) let publicPaths: [PublicPath]

    /// All private paths of this account.
    access(all) let privatePaths: [PrivatePath]

    /// All storage paths of this account.
    access(all) let storagePaths: [StoragePath]

    /// Saves the given object into the account's storage at the given path.
    ///
    /// Resources are moved into storage, and structures are copied.
    ///
    /// If there is already an object stored under the given path, the program aborts.
    ///
    /// The path must be a storage path, i.e., only the domain `storage` is allowed.
    access(all) fun save<T: Storable>(_ value: T, to: StoragePath)

    /// Reads the type of an object from the account's storage which is stored under the given path,
    /// or nil if no object is stored under the given path.
    ///
    /// If there is an object stored, the type of the object is returned without modifying the stored object.
    ///
    /// The path must be a storage path, i.e., only the domain `storage` is allowed.
    access(all) view fun type(at path: StoragePath): Type?

    /// Loads an object from the account's storage which is stored under the given path,
    /// or nil if no object is stored under the given path.
    ///
    /// If there is an object stored,
    /// the stored resource or structure is moved out of storage and returned as an optional.
    ///
    /// When the function returns, the storage no longer contains an object under the given path.
    ///
    /// The given type must be a supertype of the type of the loaded object.
    /// If it is not, the function panics.
    ///
    /// The given type must not necessarily be exactly the same as the type of the loaded object.
    ///
    /// The path must be a storage path, i.e., only the domain `storage` is allowed.
    access(all) fun load<T: Storable>(from: StoragePath): T?

    /// Returns a copy of a structure stored in account storage under the given path,
    /// without removing it from storage,
    /// or nil if no object is stored under the given path.
    ///
    /// If there is a structure stored, it is copied.
    /// The structure stays stored in storage after the function returns.
    ///
    /// The given type must be a supertype of the type of the copied structure.
    /// If it is not, the function panics.
    ///
    /// The given type must not necessarily be exactly the same as the type of the copied structure.
    ///
    /// The path must be a storage path, i.e., only the domain `storage` is allowed.
    access(all) fun copy<T: AnyStruct>(from: StoragePath): T?

    /// Returns a reference to an object in storage without removing it from storage.
    ///
    /// If no object is stored under the given path, the function returns nil.
    /// If there is an object stored, a reference is returned as an optional,
    /// provided it can be borrowed using the given type.
    /// If the stored object cannot be borrowed using the given type, the function panics.
    ///
    /// The given type must not necessarily be exactly the same as the type of the borrowed object.
    ///
    /// The path must be a storage path, i.e., only the domain `storage` is allowed
    access(all) fun borrow<T: &Any>(from: StoragePath): T?

    /// **DEPRECATED**: Instead, use `capabilities.storage.issue`, and `capabilities.publish` if the path is public.
    ///
    /// Creates a capability at the given public or private path,
    /// which targets the given public, private, or storage path.
    ///
    /// The target path leads to the object that will provide the functionality defined by this capability.
    ///
    /// The given type defines how the capability can be borrowed, i.e., how the stored value can be accessed.
    ///
    /// Returns nil if a link for the given capability path already exists, or the newly created capability if not.
    ///
    /// It is not necessary for the target path to lead to a valid object; the target path could be empty,
    /// or could lead to an object which does not provide the necessary type interface:
    /// The link function does **not** check if the target path is valid/exists at the time the capability is created
    /// and does **not** check if the target value conforms to the given type.
    ///
    /// The link is latent.
    ///
    /// The target value might be stored after the link is created,
    /// and the target value might be moved out after the link has been created.
    access(all) fun link<T: &Any>(_ newCapabilityPath: CapabilityPath, target: Path): Capability<T>?

    /// **DEPRECATED**: Use `capabilities.account.issue` instead.
    ///
    /// Creates a capability at the given public or private path which targets this account.
    ///
    /// Returns nil if a link for the given capability path already exists, or the newly created capability if not.
    access(all) fun linkAccount(_ newCapabilityPath: PrivatePath): Capability<&AuthAccount>?

    /// **DEPRECATED**: Use `capabilities.get` instead.
    ///
    /// Returns the capability at the given private or public path.
    access(all) fun getCapability<T: &Any>(_ path: CapabilityPath): Capability<T>

    /// **DEPRECATED**: Use `capabilities.storage.getController` and `StorageCapabilityController.target()`.
    ///
    /// Returns the target path of the capability at the given public or private path,
    /// or nil if there exists no capability at the given path.
    access(all) fun getLinkTarget(_ path: CapabilityPath): Path?

    /// **DEPRECATED**: Use `capabilities.unpublish` instead if the path is public.
    ///
    /// Removes the capability at the given public or private path.
    access(all) fun unlink(_ path: CapabilityPath)

    /// Iterate over all the public paths of an account,
    /// passing each path and type in turn to the provided callback function.
    ///
    /// The callback function takes two arguments:
    ///   1. The path of the stored object
    ///   2. The runtime type of that object
    ///
    /// Iteration is stopped early if the callback function returns `false`.
    ///
    /// The order of iteration is undefined.
    ///
    /// If an object is stored under a new public path,
    /// or an existing object is removed from a public path,
    /// then the callback must stop iteration by returning false.
    /// Otherwise, iteration aborts.
    ///
    access(all) fun forEachPublic(_ function: fun(PublicPath, Type): Bool)

    /// Iterate over all the private paths of an account,
    /// passing each path and type in turn to the provided callback function.
    ///
    /// The callback function takes two arguments:
    ///   1. The path of the stored object
    ///   2. The runtime type of that object
    ///
    /// Iteration is stopped early if the callback function returns `false`.
    ///

    /// The order of iteration is undefined.
    ///
    /// If an object is stored under a new private path,
    /// or an existing object is removed from a private path,
    /// then the callback must stop iteration by returning false.
    /// Otherwise, iteration aborts.
    access(all) fun forEachPrivate(_ function: fun(PrivatePath, Type): Bool)

    /// Iterate over all the stored paths of an account,
    /// passing each path and type in turn to the provided callback function.
    ///
    /// The callback function takes two arguments:
    ///   1. The path of the stored object
    ///   2. The runtime type of that object
    ///
    /// Iteration is stopped early if the callback function returns `false`.
    ///

    /// If an object is stored under a new storage path,
    /// or an existing object is removed from a storage path,
    /// then the callback must stop iteration by returning false.
    /// Otherwise, iteration aborts.
    access(all) fun forEachStored(_ function: fun(StoragePath, Type): Bool)

    access(all) struct Contracts {

        /// The names of all contracts deployed in the account.
        access(all) let names: [String]

        /// Adds the given contract to the account.
        ///
        /// The `code` parameter is the UTF-8 encoded representation of the source code.
        /// The code must contain exactly one contract or contract interface,
        /// which must have the same name as the `name` parameter.
        ///
        /// All additional arguments that are given are passed further to the initializer
        /// of the contract that is being deployed.
        ///
        /// The function fails if a contract/contract interface with the given name already exists in the account,
        /// if the given code does not declare exactly one contract or contract interface,
        /// or if the given name does not match the name of the contract/contract interface declaration in the code.
        ///
        /// Returns the deployed contract.
        access(all) fun add(
            name: String,
            code: [UInt8]
        ): DeployedContract

        /// **Experimental**
        ///
        /// Updates the code for the contract/contract interface in the account.
        ///
        /// The `code` parameter is the UTF-8 encoded representation of the source code.
        /// The code must contain exactly one contract or contract interface,
        /// which must have the same name as the `name` parameter.
        ///
        /// Does **not** run the initializer of the contract/contract interface again.
        /// The contract instance in the world state stays as is.
        ///
        /// Fails if no contract/contract interface with the given name exists in the account,
        /// if the given code does not declare exactly one contract or contract interface,
        /// or if the given name does not match the name of the contract/contract interface declaration in the code.
        ///
        /// Returns the deployed contract for the updated contract.
        access(all) fun update__experimental(name: String, code: [UInt8]): DeployedContract

        /// Returns the deployed contract for the contract/contract interface with the given name in the account, if any.
        ///
        /// Returns nil if no contract/contract interface with the given name exists in the account.
        access(all) fun get(name: String): DeployedContract?

        /// Removes the contract/contract interface from the account which has the given name, if any.
        ///
        /// Returns the removed deployed contract, if any.
        ///
        /// Returns nil if no contract/contract interface with the given name exists in the account.
        access(all) fun remove(name: String): DeployedContract?

        /// Returns a reference of the given type to the contract with the given name in the account, if any.
        ///
        /// Returns nil if no contract with the given name exists in the account,
        /// or if the contract does not conform to the given type.
        access(all) fun borrow<T: &Any>(name: String): T?
    }

    access(all) struct Keys {

        /// Adds a new key with the given hashing algorithm and a weight.
        ///
        /// Returns the added key.
        access(all) fun add(
            publicKey: PublicKey,
            hashAlgorithm: HashAlgorithm,
            weight: UFix64
        ): AccountKey

        /// Returns the key at the given index, if it exists, or nil otherwise.
        ///
        /// Revoked keys are always returned, but they have `isRevoked` field set to true.
        access(all) fun get(keyIndex: Int): AccountKey?

        /// Marks the key at the given index revoked, but does not delete it.
        ///
        /// Returns the revoked key if it exists, or nil otherwise.
        access(all) fun revoke(keyIndex: Int): AccountKey?

        /// Iterate over all unrevoked keys in this account,
        /// passing each key in turn to the provided function.
        ///
        /// Iteration is stopped early if the function returns `false`.
        ///
        /// The order of iteration is undefined.
        access(all) fun forEach(_ function: fun(AccountKey): Bool)

        /// The total number of unrevoked keys in this account.
        access(all) let count: UInt64
    }

    access(all) struct Inbox {

        /// Publishes a new Capability under the given name,
        /// to be claimed by the specified recipient.
        access(all) fun publish(_ value: Capability, name: String, recipient: Address)

        /// Unpublishes a Capability previously published by this account.
        ///
        /// Returns `nil` if no Capability is published under the given name.
        ///
        /// Errors if the Capability under that name does not match the provided type.
        access(all) fun unpublish<T: &Any>(_ name: String): Capability<T>?

        /// Claims a Capability previously published by the specified provider.
        ///
        /// Returns `nil` if no Capability is published under the given name,
        /// or if this account is not its intended recipient.
        ///
        /// Errors if the Capability under that name does not match the provided type.
        access(all) fun claim<T: &Any>(_ name: String, provider: Address): Capability<T>?
    }

    access(all) struct Capabilities {

        /// The storage capabilities of the account.
        access(all) let storage: AuthAccount.StorageCapabilities

        /// The account capabilities of the account.
        access(all) let account: AuthAccount.AccountCapabilities

        /// Returns the capability at the given public path.
        /// Returns nil if the capability does not exist,
        /// or if the given type is not a supertype of the capability's borrow type.
        access(all) fun get<T: &Any>(_ path: PublicPath): Capability<T>?

        /// Borrows the capability at the given public path.
        /// Returns nil if the capability does not exist, or cannot be borrowed using the given type.
        /// The function is equivalent to `get(path)?.borrow()`.
        access(all) fun borrow<T: &Any>(_ path: PublicPath): T?

        /// Publish the capability at the given public path.
        ///
        /// If there is already a capability published under the given path, the program aborts.
        ///
        /// The path must be a public path, i.e., only the domain `public` is allowed.
        access(all) fun publish(_ capability: Capability, at: PublicPath)

        /// Unpublish the capability published at the given path.
        ///
        /// Returns the capability if one was published at the path.
        /// Returns nil if no capability was published at the path.
        access(all) fun unpublish(_ path: PublicPath): Capability?

        /// **DEPRECATED**: This function only exists temporarily to aid in the migration of links.
        /// This function will not be part of the final Capability Controller API.
        ///
        /// Migrates the link at the given path to a capability controller.
        /// Returns the capability ID of the newly issued controller.
        /// Returns nil if the migration fails,
        /// e.g. when the path does not lead to a storage path.
        ///
        /// Does not migrate intermediate links of the chain.
        ///
        /// Returns the ID of the issued capability controller, if any.
        /// Returns nil if migration fails.
        access(all) fun migrateLink(_ newCapabilityPath: CapabilityPath): UInt64?
    }

    access(all) struct StorageCapabilities {

        /// Get the storage capability controller for the capability with the specified ID.
        ///
        /// Returns nil if the ID does not reference an existing storage capability.
        access(all) fun getController(byCapabilityID: UInt64): &StorageCapabilityController?

        /// Get all storage capability controllers for capabilities that target this storage path
        access(all) fun getControllers(forPath: StoragePath): [&StorageCapabilityController]

        /// Iterate over all storage capability controllers for capabilities that target this storage path,
        /// passing a reference to each controller to the provided callback function.
        ///
        /// Iteration is stopped early if the callback function returns `false`.
        ///
        /// If a new storage capability controller is issued for the path,
        /// an existing storage capability controller for the path is deleted,
        /// or a storage capability controller is retargeted from or to the path,
        /// then the callback must stop iteration by returning false.
        /// Otherwise, iteration aborts.
        access(all) fun forEachController(forPath: StoragePath, _ function: fun(&StorageCapabilityController): Bool)

        /// Issue/create a new storage capability.
        access(all) fun issue<T: &Any>(_ path: StoragePath): Capability<T>
    }

    access(all) struct AccountCapabilities {
        /// Get capability controller for capability with the specified ID.
        ///
        /// Returns nil if the ID does not reference an existing account capability.
        access(all) fun getController(byCapabilityID: UInt64): &AccountCapabilityController?

        /// Get all capability controllers for all account capabilities.
        access(all) fun getControllers(): [&AccountCapabilityController]

        /// Iterate over all account capability controllers for all account capabilities,
        /// passing a reference to each controller to the provided callback function.
        ///
        /// Iteration is stopped early if the callback function returns `false`.
        ///
        /// If a new account capability controller is issued for the account,
        /// or an existing account capability controller for the account is deleted,
        /// then the callback must stop iteration by returning false.
        /// Otherwise, iteration aborts.
        access(all) fun forEachController(_ function: fun(&AccountCapabilityController): Bool)

        /// Issue/create a new account capability.
        access(all) fun issue<T: &AuthAccount{}>(): Capability<T>
    }
}
