(function() {
    'use strict';

    var SECTIONS = {
        'dashboard': 'sections/dashboard.html',
        'system': 'sections/system.html'
    };

    var currentSection = 'dashboard';
    var sectionCache = {};

    window.showErrorAndLogout = function(message, detail) {
        if (typeof Layer8MAuth !== 'undefined') {
            Layer8MAuth.showErrorAndLogout(message, detail);
        } else {
            alert(message + (detail ? '\n\n' + detail : ''));
            window.location.href = '/l8ui/login/';
        }
    };

    window.MobileApp = {
        async init() {
            if (!Layer8MAuth.requireAuth()) return;

            await Layer8MConfig.load();
            await Layer8DConfig.load();

            this.updateUserInfo();

            var token = Layer8MAuth.getBearerToken();

            try {
                var permResp = await fetch('/permissions', {
                    headers: { 'Authorization': 'Bearer ' + token, 'Content-Type': 'application/json' }
                });
                if (permResp.ok) {
                    window.Layer8DPermissions = await permResp.json();
                }
            } catch (e) { console.warn('Failed to load permissions:', e); }

            this.initSidebar();

            document.getElementById('refresh-btn')?.addEventListener('click', function() {
                MobileApp.loadSection(currentSection, true);
            });

            var hash = window.location.hash.slice(1);
            var section = SECTIONS[hash] ? hash : 'dashboard';
            await this.loadSection(section);

            window.addEventListener('hashchange', function() {
                var newSection = window.location.hash.slice(1);
                if (SECTIONS[newSection] && newSection !== currentSection) {
                    MobileApp.loadSection(newSection);
                }
            });
        },

        updateUserInfo() {
            var username = Layer8MAuth.getUsername();
            var initial = username.charAt(0).toUpperCase();
            document.getElementById('user-name').textContent = username;
            document.getElementById('user-avatar').textContent = initial;
        },

        initSidebar() {
            var menuToggle = document.getElementById('menu-toggle');
            var overlay = document.getElementById('sidebar-overlay');

            if (menuToggle) menuToggle.addEventListener('click', function() { MobileApp.openSidebar(); });
            if (overlay) overlay.addEventListener('click', function() { MobileApp.closeSidebar(); });

            document.querySelectorAll('.sidebar-item[data-section]').forEach(function(item) {
                item.addEventListener('click', async function(e) {
                    e.preventDefault();
                    var section = item.dataset.section;
                    var module = item.dataset.module;
                    MobileApp.closeSidebar();
                    await MobileApp.loadSection(section);
                    if (module && window.Layer8MNav) {
                        Layer8MNav.navigateToModule(module);
                    }
                });
            });
        },

        openSidebar() {
            document.getElementById('sidebar')?.classList.add('open');
            document.getElementById('sidebar-overlay')?.classList.add('visible');
            document.body.style.overflow = 'hidden';
        },

        closeSidebar() {
            document.getElementById('sidebar')?.classList.remove('open');
            document.getElementById('sidebar-overlay')?.classList.remove('visible');
            document.body.style.overflow = '';
        },

        async loadSection(section, forceReload) {
            if (section !== 'dashboard' && window.LAYER8M_NAV_CONFIG && LAYER8M_NAV_CONFIG[section]) {
                await this._loadDashboardForModule(section, forceReload);
                return;
            }

            var sectionUrl = SECTIONS[section];
            if (!sectionUrl) {
                console.error('Unknown section:', section);
                return;
            }

            this.updateNavState(section);

            var contentArea = document.getElementById('content-area');
            if (!contentArea) return;

            contentArea.style.opacity = '0.5';

            try {
                if (!forceReload && sectionCache[section]) {
                    contentArea.innerHTML = sectionCache[section];
                } else {
                    var response = await fetch(sectionUrl + '?t=' + Date.now());
                    if (!response.ok) throw new Error('Failed to load section');
                    var html = await response.text();
                    sectionCache[section] = html;
                    contentArea.innerHTML = html;
                }

                this.executeScripts(contentArea);
                this.initSection(section);
                currentSection = section;
                window.location.hash = section;
                contentArea.scrollTop = 0;
            } catch (error) {
                console.error('Error loading section:', error);
                contentArea.innerHTML =
                    '<div class="empty-state">' +
                    '<span class="empty-state-icon">&#x26A0;</span>' +
                    '<h4 class="empty-state-title">Failed to load</h4>' +
                    '<p class="empty-state-message">Please try again</p>' +
                    '<button class="btn btn-primary" onclick="MobileApp.loadSection(\'' + section + '\', true)">Retry</button>' +
                    '</div>';
            }

            contentArea.style.opacity = '1';
        },

        async _loadDashboardForModule(moduleKey, forceReload) {
            this.updateNavState(moduleKey);

            var contentArea = document.getElementById('content-area');
            if (!contentArea) return;

            contentArea.style.opacity = '0.5';

            try {
                if (!forceReload && sectionCache['dashboard']) {
                    contentArea.innerHTML = sectionCache['dashboard'];
                } else {
                    var response = await fetch(SECTIONS['dashboard'] + '?t=' + Date.now());
                    if (!response.ok) throw new Error('Failed to load dashboard');
                    var html = await response.text();
                    sectionCache['dashboard'] = html;
                    contentArea.innerHTML = html;
                }

                this.executeScripts(contentArea);
                this.initSection('dashboard');
                Layer8MNav.navigateToModule(moduleKey);

                currentSection = moduleKey;
                window.location.hash = moduleKey;
                contentArea.scrollTop = 0;
            } catch (error) {
                console.error('Error loading module:', error);
            }

            contentArea.style.opacity = '1';
        },

        updateNavState(section) {
            document.querySelectorAll('.sidebar-item').forEach(function(item) {
                item.classList.remove('active');
                if (item.dataset.section === section) {
                    item.classList.add('active');
                }
            });
        },

        executeScripts(container) {
            var scripts = container.querySelectorAll('script');
            scripts.forEach(function(oldScript) {
                var newScript = document.createElement('script');
                Array.from(oldScript.attributes).forEach(function(attr) {
                    newScript.setAttribute(attr.name, attr.value);
                });
                newScript.textContent = oldScript.textContent;
                oldScript.parentNode.replaceChild(newScript, oldScript);
            });
        },

        initSection(section) {
            var initFunctions = {
                'dashboard': 'initMobileDashboard',
                'system': 'initMobileSystem'
            };

            var initFn = initFunctions[section];
            if (initFn && typeof window[initFn] === 'function') {
                window[initFn]();
            }
        },

        getCurrentSection() {
            return currentSection;
        },

        logout() {
            Layer8MAuth.logout();
        }
    };

    if (document.readyState === 'loading') {
        document.addEventListener('DOMContentLoaded', function() { MobileApp.init(); });
    } else {
        MobileApp.init();
    }
})();
