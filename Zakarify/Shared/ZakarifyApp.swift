//
//  ZakarifyApp.swift
//  Shared
//
//  Created by Zakariah Music on 8/12/21.
//

import SwiftUI

@main
struct ZakarifyApp: App {
    let persistenceController = PersistenceController.shared

    var body: some Scene {
        WindowGroup {
            ContentView()
                .environment(\.managedObjectContext, persistenceController.container.viewContext)
        }
    }
}
