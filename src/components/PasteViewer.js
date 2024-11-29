import React, { useState, useEffect } from 'react';
import { getPaste } from '../services/api';
import { useParams, Link } from 'react-router-dom';
import { Container, Alert, Button } from 'react-bootstrap';

const PasteViewer = () => {
    const [content, setContent] = useState('Loading...');
    const { id } = useParams();

    useEffect(() => {
        const fetchContent = async () => {
            const data = await getPaste(id);
            if (data.content) {
                setContent(data.content);
            } else {
                setContent('Paste not found.');
            }
        };
        fetchContent();
    }, [id]);

    return (
        <Container className="mt-3">
            {content === 'Paste not found.' ? (
                <Alert variant="danger">{content}</Alert>
            ) : (
                <pre>{content}</pre>
            )}
        </Container>
    );
};

export default PasteViewer;