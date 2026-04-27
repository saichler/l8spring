/*
© 2025 Sharon Aicler (saichler@gmail.com)

Layer 8 Ecosystem is licensed under the Apache License, Version 2.0.
You may obtain a copy of the License at:

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

function getAuthHeaders() {
    var bearerToken = sessionStorage.getItem('bearerToken');
    return {
        'Authorization': bearerToken ? 'Bearer ' + bearerToken : '',
        'Content-Type': 'application/json'
    };
}

function logout() {
    sessionStorage.removeItem('bearerToken');
    localStorage.removeItem('bearerToken');
    window.location.href = 'l8ui/login/index.html';
}

function showErrorAndLogout(message, detail) {
    sessionStorage.removeItem('bearerToken');
    localStorage.removeItem('bearerToken');
    if (typeof Layer8DPopup !== 'undefined') {
        Layer8DPopup.show({
            title: 'Session Error',
            content: '<div style="padding:16px;"><p style="margin-bottom:12px;font-size:15px;">' +
                (typeof Layer8DUtils !== 'undefined' ? Layer8DUtils.escapeHtml(message) : message) + '</p>' +
                (detail ? '<pre style="background:var(--layer8d-bg-light);padding:12px;border-radius:6px;font-size:12px;max-height:200px;overflow:auto;white-space:pre-wrap;word-break:break-word;">' +
                (typeof Layer8DUtils !== 'undefined' ? Layer8DUtils.escapeHtml(detail) : detail) + '</pre>' : '') +
                '</div>',
            size: 'medium',
            showFooter: true,
            saveButtonText: 'Go to Login',
            showCancelButton: false,
            onSave: function() {
                Layer8DPopup.close();
                window.location.href = 'l8ui/login/index.html';
            }
        });
    } else {
        alert(message + (detail ? '\n\n' + detail : ''));
        window.location.href = 'l8ui/login/index.html';
    }
}

function loadSection(sectionName) {
    var path = window.sections[sectionName];
    if (!path) return;

    var content = document.getElementById('content-area');
    if (!content) return;

    document.querySelectorAll('.nav-link').forEach(function(el) {
        el.classList.remove('active');
        if (el.getAttribute('data-section') === sectionName) el.classList.add('active');
    });

    fetch(path)
        .then(function(r) { return r.text(); })
        .then(function(html) {
            content.innerHTML = html;
            var scripts = content.querySelectorAll('script');
            scripts.forEach(function(oldScript) {
                var newScript = document.createElement('script');
                Array.from(oldScript.attributes).forEach(function(attr) {
                    newScript.setAttribute(attr.name, attr.value);
                });
                newScript.textContent = oldScript.textContent;
                oldScript.parentNode.replaceChild(newScript, oldScript);
            });
            var init = window.sectionInitializers[sectionName];
            if (init) init();
        })
        .catch(function(e) {
            console.error('Failed to load section:', e);
        });
}

document.addEventListener('DOMContentLoaded', async function() {
    if (typeof Layer8DConfig !== 'undefined') {
        await Layer8DConfig.load();
    }

    var bearerToken = sessionStorage.getItem('bearerToken');
    if (!bearerToken) {
        window.location.href = 'l8ui/login/index.html';
        return;
    }

    localStorage.setItem('bearerToken', bearerToken);
    window.bearerToken = bearerToken;

    var usernameEl = document.querySelector('.username');
    if (usernameEl) {
        usernameEl.textContent = sessionStorage.getItem('currentUser') || 'Admin';
    }

    try {
        var permResp = await fetch('/permissions', {
            headers: { 'Authorization': 'Bearer ' + bearerToken, 'Content-Type': 'application/json' }
        });
        if (permResp.ok) {
            window.Layer8DPermissions = await permResp.json();
        }
    } catch (e) { console.warn('Failed to load permissions:', e); }

    loadSection('dashboard');

    var navLinks = document.querySelectorAll('.nav-link');
    navLinks.forEach(function(link) {
        link.addEventListener('click', function(e) {
            e.preventDefault();
            navLinks.forEach(function(l) { l.classList.remove('active'); });
            this.classList.add('active');
            var section = this.getAttribute('data-section');
            loadSection(section);
        });
    });
});
