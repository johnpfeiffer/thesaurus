import React, { useState } from 'react';
import './Thesaurus.css';
import thesaurus from '../data/thesaurus';

function Thesaurus() {
  const [inputValue, setInputValue] = useState('');
  const [suggestion, setSuggestion] = useState('');

  const handleInputChange = (event) => {
    setInputValue(event.target.value);
  };

  const handleButtonClick = () => {
    const suggestion = thesaurus[inputValue.toLowerCase()] || 'No suggestion available';
    setSuggestion(suggestion);
  };

  return (
    <div className="Thesaurus">
      <input
        type="text"
        value={inputValue}
        onChange={handleInputChange}
        maxLength="2"
        placeholder="Enter a two-letter word"
      />
      <button onClick={handleButtonClick}>Get Suggestion</button>
      <p>Suggestion: {suggestion}</p>
    </div>
  );
}

export default Thesaurus;
