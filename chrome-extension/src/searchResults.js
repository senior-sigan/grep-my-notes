(async () => {
  function noteBlock(note) {
    return `
      <div class="note_block">
        <a class="note_link" target="_blank" href="${note.link}">
          <h3>${note.title}</h3>
        </a>
        <article class="note_body">${note.slug}</article>
      </div>
    `;
  }

  async function renderResults(results) {
    const notesBlock = document.getElementById('notes-list');
    notesBlock.innerHTML = '';

    const data = await results.json();
    if (!data || !data.entries) {
      return;
    }

    data.entries.forEach((entry) => {
      notesBlock.insertAdjacentHTML(
        'beforeend',
        noteBlock({
          link: `vscode://file${entry.file}`,
          title: entry.title,
          slug: entry.slug,
        }),
      );
    });
  }

  function renderError(error) {
    const notesBlock = document.getElementById('notesCards');
    notesBlock.insertAdjacentHTML(
      'afterbegin',
      noteBlock({
        link: '',
        date: '',
        title: 'ERROR',
        snippet: error,
      }),
    );
  }

  async function getSearchQuery() {
    const tabs = await chrome.tabs.query({ active: true, currentWindow: true });
    const url = new URL(tabs[0].url);
    const params = new URLSearchParams(url.search);

    return params.get('q');
  }

  const q = await getSearchQuery();
  try {
    const result = await fetch(`http://localhost:3000/?query=${q}`);
    renderResults(result);
  } catch (err) {
    renderError(err);
  }
})();
