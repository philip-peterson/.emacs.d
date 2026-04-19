(require 'package)
(add-to-list 'package-archives '("melpa" . "https://melpa.org/packages/") t)
(package-initialize)

;; make sure m- works
(setq mac-option-modifier 'meta)

(use-package meow
  :ensure t
  :init
  (defun meow-setup ())
  :config
  (meow-setup)
  (meow-global-mode 1))


(use-package treemacs
  :ensure t
  :defer t
  :custom (treemacs-position 'right)
  ;; :hook ((prog-mode . treemacs))
  )

;; ctrl-_ for undo
;;(global-set-key (kbd "C-Z") undo)

(custom-set-variables
 ;; custom-set-variables was added by Custom.
 ;; If you edit it by hand, you could mess it up, so be careful.
 ;; Your init file should contain only one such instance.
 ;; If there is more than one, they won't work right.
 '(package-selected-packages
   '(lsp-mode meow vertico treemacs orderless lsp-ui go-mode company)))
(custom-set-faces
 ;; custom-set-faces was added by Custom.
 ;; If you edit it by hand, you could mess it up, so be careful.
 ;; Your init file should contain only one such instance.
 ;; If there is more than one, they won't work right.
 )

(use-package orderless
  :ensure t
  :custom
  (completion-styles '(orderless basic))
  (completion-category-overrides '((file (styles basic partial-completion)))))

(use-package company
  :ensure t
  :diminish company-mode
  :hook prog-mode)

(use-package vertico
  :ensure t
  ;; (vertico-scroll-margin 0) ;; Different scroll margin
  ;; (vertico-count 20) ;; Show more candidates
  ;; (vertico-resize t) ;; Grow and shrink the Vertico minibuffer
  ;; (vertico-cycle t) ;; Enable cycling for `vertico-next/previous'
  :init
  (vertico-mode))


;; Install go mode

(use-package lsp-mode
  :ensure t
  :hook ((go-mode . lsp))
  :commands lsp)

(use-package go-mode
  :ensure t
  :mode "\\.go\\'")

(use-package lsp-ui :ensure t)
