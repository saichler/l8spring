(function() {
    'use strict';

    var bearerToken = sessionStorage.getItem('bearerToken');
    if (!bearerToken) {
        window.location.href = '/login.html';
        return;
    }

    window.getAuthHeaders = function() {
        return {
            'Authorization': 'Bearer ' + bearerToken,
            'Content-Type': 'application/json'
        };
    };

    window.logout = function() {
        sessionStorage.removeItem('bearerToken');
        sessionStorage.removeItem('username');
        window.location.href = '/login.html';
    };

    var currentSection = null;

    async function initApp() {
        try {
            await Layer8DConfig.load();
        } catch (e) {
            console.error('Failed to load config:', e);
        }

        try {
            var permResp = await fetch('/permissions', {
                headers: getAuthHeaders()
            });
            if (permResp.ok) {
                window.Layer8DPermissions = await permResp.json();
            }
        } catch (e) { console.warn('Failed to load permissions:', e); }

        loadSidebar();
        loadDefaultSection();
    }

    function loadSidebar() {
        var sidebar = document.getElementById('sidebar-nav');
        if (!sidebar) return;

        var links = [
            { section: 'dashboard', label: 'Dashboard', icon: '<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect x="3" y="3" width="7" height="7"/><rect x="14" y="3" width="7" height="7"/><rect x="14" y="14" width="7" height="7"/><rect x="3" y="14" width="7" height="7"/></svg>' },
            { section: 'marketplace', label: 'Marketplace', icon: '<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="9" cy="21" r="1"/><circle cx="20" cy="21" r="1"/><path d="M1 1h4l2.68 13.39a2 2 0 0 0 2 1.61h9.72a2 2 0 0 0 2-1.61L23 6H6"/></svg>' },
            { section: 'system', label: 'System', icon: '<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="3"/><path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 0 1 0 2.83 2 2 0 0 1-2.83 0l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-2 2 2 2 0 0 1-2-2v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 0 1-2.83 0 2 2 0 0 1 0-2.83l.06-.06A1.65 1.65 0 0 0 4.68 15a1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1-2-2 2 2 0 0 1 2-2h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 0 1 0-2.83 2 2 0 0 1 2.83 0l.06.06A1.65 1.65 0 0 0 9 4.68a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 2-2 2 2 0 0 1 2 2v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 0 1 2.83 0 2 2 0 0 1 0 2.83l-.06.06A1.65 1.65 0 0 0 19.4 9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 2 2 2 2 0 0 1-2 2h-.09a1.65 1.65 0 0 0-1.51 1z"/></svg>' }
        ];

        links.forEach(function(link) {
            var a = document.createElement('a');
            a.href = '#' + link.section;
            a.className = 'sidebar-item';
            a.setAttribute('data-section', link.section);
            if (link.icon) {
                var iconSpan = document.createElement('span');
                iconSpan.className = 'sidebar-icon';
                iconSpan.innerHTML = link.icon;
                a.appendChild(iconSpan);
            }
            var labelSpan = document.createElement('span');
            labelSpan.textContent = link.label;
            a.appendChild(labelSpan);
            a.addEventListener('click', function(e) {
                e.preventDefault();
                loadSection(link.section);
            });
            sidebar.appendChild(a);
        });
    }

    function loadSection(sectionName) {
        var path = window.sections[sectionName];
        if (!path) return;

        var content = document.getElementById('main-content');
        if (!content) return;

        document.querySelectorAll('.sidebar-item').forEach(function(el) {
            el.classList.remove('active');
            if (el.getAttribute('data-section') === sectionName) el.classList.add('active');
        });

        currentSection = sectionName;

        fetch(path)
            .then(function(r) { return r.text(); })
            .then(function(html) {
                content.innerHTML = html;
                executeScripts(content);
                var init = window.sectionInitializers[sectionName];
                if (init) init();
            })
            .catch(function(e) {
                console.error('Failed to load section:', e);
            });
    }

    function executeScripts(container) {
        var scripts = container.querySelectorAll('script');
        scripts.forEach(function(oldScript) {
            var newScript = document.createElement('script');
            Array.from(oldScript.attributes).forEach(function(attr) {
                newScript.setAttribute(attr.name, attr.value);
            });
            newScript.textContent = oldScript.textContent;
            oldScript.parentNode.replaceChild(newScript, oldScript);
        });
    }

    function loadDefaultSection() {
        var hash = window.location.hash.slice(1);
        var section = window.sections[hash] ? hash : 'dashboard';
        loadSection(section);
    }

    window.addEventListener('hashchange', function() {
        var hash = window.location.hash.slice(1);
        if (window.sections[hash] && hash !== currentSection) {
            loadSection(hash);
        }
    });

    window.loadSection = loadSection;

    initApp();
})();
