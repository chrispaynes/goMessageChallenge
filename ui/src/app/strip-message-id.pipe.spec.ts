import { StripMessageIdPipe } from './strip-message-id.pipe';

describe('StripMessageIdPipe', () => {
  it('create an instance', () => {
    const pipe = new StripMessageIdPipe();
    expect(pipe).toBeTruthy();
  });
});
