//
//  ViewController.swift
//  iOS14NewFeature
//
//  Created by Hiroaki Tomiyoshi on 2020/11/02.
//

import UIKit
import StoreKit

class ViewController: UIViewController {

    override func viewDidLoad() {
        super.viewDidLoad()
        // Do any additional setup after loading the view.
    }

    @IBAction func skOverlayButtonTapped(_ sender: Any) {
        if let scene = view.window?.windowScene {
            let config = SKOverlay.AppConfiguration(appIdentifier: "376101648", position: .bottom) // App Store アプリのID
            let overlay = SKOverlay(configuration: config)
            overlay.present(in: scene)
        }
    }
}

