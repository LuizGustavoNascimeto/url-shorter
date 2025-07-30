import { useState } from 'react';
import './App.css';

function App() {
  const [originalUrl, setOriginalUrl] = useState('');
  const [shortUrl, setShortUrl] = useState('');
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    setShortUrl('');
    setLoading(true);
    try {
      const response = await fetch('/api/shorten', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ original_url: originalUrl })
      });
      const data = await response.json();
      if (response.ok) {
        // Tenta pegar o short_url direto, se não, extrai do message
        if (data.short_url) {
          setShortUrl(`/${data.short_url}`);
        } else if (data.message) {
          const match = data.message.match(/para (\w+)$/);
          setShortUrl(match ? `/${match[1]}` : '');
        }
      } else {
        setError(data.error || 'Erro ao encurtar a URL.');
      }
    } catch {
      setError('Erro de conexão com o servidor.');
    }
    setLoading(false);
  };

  return (
    <div className="container">
      <h1>URL Shortener</h1>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          placeholder="Digite a URL aqui"
          value={originalUrl}
          onChange={e => setOriginalUrl(e.target.value)}
          disabled={loading}
        />
        <button type="submit" disabled={loading}>
          {loading ? 'Encurtando...' : 'Encurtar URL'}
        </button>
      </form>
      {shortUrl && (
        <div className="result">
          URL encurtada: <a href={"http://localhost:8080" + shortUrl} target="_blank" rel="noopener noreferrer">{"http://localhost:8080" + shortUrl}</a>
        </div>
      )}
      {error && <div className="error">{error}</div>}
    </div>
  );
}

export default App;
