import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { createPaste } from '../services/api';
import { Form, Button, Container } from 'react-bootstrap';

const PasteForm = () => {
    const [content, setContent] = useState('');
    const navigate = useNavigate();

    const handleSubmit = async (e) => {
        e.preventDefault();
        const data = await createPaste(content);
        if (data.id) {
            navigate(`/paste/${data.id}`);
        } else {
            alert('Failed to create paste.');
        }
    };

    return (
        <Container>
            <Form onSubmit={handleSubmit}>
                <Form.Group controlId="pasteContent">
                    <Form.Label></Form.Label>
                    <Form.Control
                        as="textarea"
                        rows={10}
                        value={content}
                        onChange={(e) => setContent(e.target.value)}
                        placeholder="Enter your text here..."
                        required
                    />
                </Form.Group>
                <Button variant="primary" type="submit" className="mt-2">
                    Create Paste
                </Button>
            </Form>
        </Container>
    );
};

export default PasteForm;